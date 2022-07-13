package service

import (
	"context"
	"fmt"
	config "graphql-golang/config"
	"graphql-golang/graph/model"
	"graphql-golang/graph/mypkg"
	db "graphql-golang/internal"
	"graphql-golang/utils"

	"github.com/dgrijalva/jwt-go"
)

type IAuthService interface {
	Signin(ctx context.Context, input *model.SigninInput) (*model.JWTResponse, error)
	Signup(ctx context.Context, input *model.SignupInput) (*model.JWTResponse, error)
	RefreshToken(ctx context.Context, refreshToken *mypkg.JWT) (*model.JWTResponse, error)
}

type AuthService struct {
	server *config.Server
}

func NewAuthService(server *config.Server) *AuthService {
	return &AuthService{
		server: server,
	}
}

func (s *AuthService) Signin(ctx context.Context, input *model.SigninInput) (*model.JWTResponse, error) {
	arg := db.LoginUserParams{
		Email: string(input.Email),
		Crypt: input.Password,
	}
	user, err := s.server.Store.LoginUser(ctx, arg)
	if err != nil {
		return nil, utils.ErrorResponse("ERROR_LOGIN_USER", err)
	}
	t, r, expt, err := s.server.GenerateJwtToken(user.ID, string(user.Role))
	if err != nil {
		return nil, utils.ErrorResponse("ERROR_TOKEN", err)
	}
	if err := s.server.StoreRefresh(ctx, r, expt, user.ID); err != nil {
		return nil, utils.ErrorResponse("ERROR_REFRESH_TOKEN", err)
	}
	response := model.JWTResponse{
		AccessToken:  mypkg.JWT(t),
		RefreshToken: mypkg.JWT(r),
		Success:      true,
	}
	return &response, nil
}

func (s *AuthService) Signup(ctx context.Context, input *model.SignupInput) (*model.JWTResponse, error) {
	if input.Password != input.ConfirmPassword {
		return nil, fmt.Errorf("password and confirm password not match")
	}
	isExist, err := s.server.Store.CheckEmailExist(ctx, string(input.Email))
	if err != nil {
		return nil, utils.ErrorResponse("ERROR_GET_MAIL", err)
	}
	if isExist {
		return nil, utils.ErrorResponse("EMAIL_ALREADY_EXIST", fmt.Errorf("email already exist"))
	}
	arg := db.SignupParams{
		Email: string(input.Email),
		Crypt: input.Password,
	}
	user, err := s.server.Store.Signup(ctx, arg)
	if err != nil {
		return nil, utils.ErrorResponse("ERROR_CREATE_USER", err)
	}
	t, r, expt, err := s.server.GenerateJwtToken(user.ID, string(user.Role))
	if err != nil {
		return nil, utils.ErrorResponse("ERROR_TOKEN", err)
	}
	if err := s.server.StoreRefresh(ctx, r, expt, user.ID); err != nil {
		return nil, utils.ErrorResponse("ERROR_REFRESH_TOKEN", err)
	}
	response := model.JWTResponse{
		AccessToken:  mypkg.JWT(t),
		RefreshToken: mypkg.JWT(r),
		Success:      true,
	}
	return &response, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken *mypkg.JWT) (*model.JWTResponse, error) {
	token, err := jwt.Parse(string(*refreshToken), func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(s.server.Config.Security.Secret))
		return b, nil
	})
	if err != nil {
		return nil, utils.ErrorResponse("TOKEN_ERROR", err)
	}
	if !token.Valid {
		return nil, utils.ErrorResponse("TOKEN_IS_NOT_VALID", fmt.Errorf("invalid format token"))
	}
	refresh, err := s.server.Store.GetRefreshToken(ctx, string(*refreshToken))
	if err != nil {
		return nil, utils.ErrorResponse("FIND_REFRESH_TOKEN", err)
	}
	t, r, expt, err := s.server.GenerateJwtToken(refresh.UserID, string(refresh.UserRole))
	if err != nil {
		return nil, utils.ErrorResponse("ERROR_TOKEN", err)
	}
	if err := s.server.StoreRefresh(ctx, r, expt, refresh.UserID); err != nil {
		return nil, utils.ErrorResponse("ERROR_REFRESH_TOKEN", err)
	}
	response := model.JWTResponse{
		AccessToken:  mypkg.JWT(t),
		RefreshToken: mypkg.JWT(r),
		Success:      true,
	}
	return &response, nil
}
