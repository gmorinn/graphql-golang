package storage

import (
	"context"
	"fmt"
	"graphql-golang/config"
	"graphql-golang/graph/model"
	"graphql-golang/graph/mypkg"
	db "graphql-golang/internal"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// UserReader reads Users from a database
type UserReader struct {
	store *config.Store
}

func (u *UserReader) GetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	usersIds := make([]string, len(keys))
	for i, key := range keys {
		usersIds[i] = key.String()
	}
	statement := fmt.Sprintf("SELECT  id, created_at, updated_at, deleted_at, email, password, name, role FROM students WHERE id IN (")
	for i, id := range usersIds {
		if i != 0 {
			statement += ","
		}
		statement += fmt.Sprintf("'%s'", id)
	}
	statement += ")"
	log.Println("==> ", statement)
	res, err := u.store.Db.Query(statement)
	if err != nil {
		panic(err)
	}
	defer res.Close()
	userById := map[string]*model.Student{}
	for res.Next() {
		user := db.Student{}
		if err := res.Scan(
			&user.ID,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.Role,
		); err != nil {
			panic(err)
		}
		userById[user.ID.String()] = &model.Student{
			ID:        mypkg.UUID(user.ID.String()),
			Email:     mypkg.Email(user.Email),
			Name:      user.Name.String,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: &user.DeletedAt.Time,
			Role:      model.UserType(user.Role),
		}
	}
	// return users in the same order requested
	output := make([]*dataloader.Result, len(keys))
	for i, userKey := range keys {
		user, ok := userById[userKey.String()]
		if ok {
			output[i] = &dataloader.Result{Data: user, Error: nil}
		} else {
			err := fmt.Errorf("user not found %s", userKey.String())
			output[i] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	UserLoader *dataloader.Loader
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(store *config.Store) *Loaders {
	// define the data loader
	userReader := &UserReader{store: store}
	loaders := &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userReader.GetUsers),
	}
	return loaders
}

// Middleware injects data loaders into the context
func Middleware(loaders *Loaders) gin.HandlerFunc {
	// return a middleware that injects the loader to the request context
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), loadersKey, loaders)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// DataloaderFor returns the dataloader for a given context
func DataloaderFor(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetUser wraps the User dataloader for efficient retrieval by user ID
func GetUser(ctx context.Context, userID mypkg.UUID) (*model.GetStudentResponse, error) {
	loaders := DataloaderFor(ctx)
	if loaders == nil {
		log.Println("no loaders")
	}
	log.Println("loaders")
	thunk := loaders.UserLoader.Load(ctx, dataloader.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	res := result.(*model.Student)
	return &model.GetStudentResponse{
		Student: res,
		Success: true,
	}, nil
}
