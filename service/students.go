package service

import (
	"context"
	config "graphql-golang/config"
	"graphql-golang/graph/model"
	"graphql-golang/graph/mypkg"
	db "graphql-golang/internal"
	"graphql-golang/utils"
	"strings"

	"github.com/google/uuid"
)

type IStudentService interface {
	UpdateStudent(ctx context.Context, input *model.UpdateStudentInput) (*model.GetStudentResponse, error)
	GetStudents(ctx context.Context, limit mypkg.NonNegativeInt, offset mypkg.NonNegativeInt) (*model.GetStudentsResponse, error)
	GetStudentByID(ctx context.Context, id mypkg.UUID) (*model.GetStudentResponse, error)
	UpdateRole(ctx context.Context, role *model.UserType, id *mypkg.UUID) (*model.GetStudentResponse, error)
}

type StudentService struct {
	server *config.Server
}

func NewStudentService(server *config.Server) *StudentService {
	return &StudentService{
		server: server,
	}
}

func (s *StudentService) UpdateStudent(ctx context.Context, input *model.UpdateStudentInput) (*model.GetStudentResponse, error) {
	var res model.Student

	var err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		id := string(input.ID)

		stud, err := q.GetStudentByID(ctx, uuid.MustParse(id))
		if err != nil {
			return err
		}

		inputEmail := string(*input.Email)

		if err := q.UpdateStudent(ctx, db.UpdateStudentParams{
			ID:    uuid.MustParse(id),
			Email: utils.UpdateString(&stud.Email, &inputEmail),
			Name:  utils.NullS(utils.UpdateString(&stud.Name.String, input.Name)),
		}); err != nil {
			return err
		}

		stud, err = q.GetStudentByID(ctx, uuid.MustParse(id))
		if err != nil {
			return err
		}

		res = model.Student{
			ID:        mypkg.UUID(stud.ID.String()),
			Email:     mypkg.Email(stud.Email),
			Name:      stud.Name.String,
			CreatedAt: stud.CreatedAt,
			UpdatedAt: stud.UpdatedAt,
			DeletedAt: &stud.DeletedAt.Time,
			Role:      model.UserType(stud.Role),
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_UPDATE_STUDENT", err)
	}

	return &model.GetStudentResponse{
		Student: &res,
		Success: true,
	}, nil
}

func (s *StudentService) GetStudents(ctx context.Context, limit mypkg.NonNegativeInt, offset mypkg.NonNegativeInt) (*model.GetStudentsResponse, error) {
	var res []*model.Student = make([]*model.Student, 0)

	var err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		studs, err := q.Liststudents(ctx, db.ListstudentsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			return err
		}

		for _, stud := range studs {
			res = append(res, &model.Student{
				ID:        mypkg.UUID(stud.ID.String()),
				Email:     mypkg.Email(stud.Email),
				Name:      stud.Name.String,
				CreatedAt: stud.CreatedAt,
				UpdatedAt: stud.UpdatedAt,
				DeletedAt: &stud.DeletedAt.Time,
				Role:      model.UserType(stud.Role),
			})
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_STUDENTS", err)
	}

	return &model.GetStudentsResponse{
		Students: res,
		Success:  true,
	}, nil
}

func (s *StudentService) GetStudentByID(ctx context.Context, id mypkg.UUID) (*model.GetStudentResponse, error) {
	var res model.Student

	var err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		id := string(id)

		stud, err := q.GetStudentByID(ctx, uuid.MustParse(id))
		if err != nil {
			return err
		}

		res = model.Student{
			ID:        mypkg.UUID(stud.ID.String()),
			Email:     mypkg.Email(stud.Email),
			Name:      stud.Name.String,
			CreatedAt: stud.CreatedAt,
			UpdatedAt: stud.UpdatedAt,
			DeletedAt: &stud.DeletedAt.Time,
			Role:      model.UserType(stud.Role),
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_STUDENT_BY_ID", err)
	}

	return &model.GetStudentResponse{
		Student: &res,
		Success: true,
	}, nil
}

func (s *StudentService) UpdateRole(ctx context.Context, role *model.UserType, id *mypkg.UUID) (*model.GetStudentResponse, error) {
	var res model.Student

	var err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		newRole := *role
		if err := q.UpdateRoleStudent(ctx, db.UpdateRoleStudentParams{
			ID:   uuid.MustParse(string(*id)),
			Role: db.Role(strings.ToLower(newRole.String())),
		}); err != nil {
			return err
		}

		stud, err := q.GetStudentByID(ctx, uuid.MustParse(string(*id)))
		if err != nil {
			return err
		}

		res = model.Student{
			ID:        mypkg.UUID(stud.ID.String()),
			Email:     mypkg.Email(stud.Email),
			Name:      stud.Name.String,
			CreatedAt: stud.CreatedAt,
			UpdatedAt: stud.UpdatedAt,
			DeletedAt: &stud.DeletedAt.Time,
			Role:      model.UserType(stud.Role),
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_UPDATE_ROLE", err)
	}

	return &model.GetStudentResponse{
		Student: &res,
		Success: true,
	}, nil
}
