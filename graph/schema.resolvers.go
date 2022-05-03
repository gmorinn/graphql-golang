package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-golang/graph/generated"
	"graphql-golang/graph/model"
	"graphql-golang/graph/mypkg"
	"graphql-golang/utils"
	"time"

	"github.com/google/uuid"
)

func (r *mutationResolver) AddOrUpdateStudent(ctx context.Context, input model.StudentInput) (*model.GetStudentResponse, error) {
	var res model.GetStudentResponse
	var user model.Student
	user.Name = input.Name
	user.Age = input.Age
	user.Gpa = input.Gpa
	user.URL = input.URL
	user.Jwt = input.Jwt
	user.Email = input.Email

	n := len(r.Resolver.StudentStore)
	if n == 0 {
		r.Resolver.StudentStore = make(map[string]model.Student)
	}

	if input.ID != nil {
		id := *input.ID
		cs, ok := r.Resolver.StudentStore[string(id)]
		if !ok {
			return nil, utils.ErrorResponse(ctx, "USER_NOT_FOUND", fmt.Errorf("user not found"))
		}
		// is_premium
		if input.IsGenius != nil {
			user.IsGenius = *input.IsGenius
		} else {
			user.IsGenius = cs.IsGenius
		}
		// role
		if input.Role != nil {
			user.Role = *input.Role
		} else {
			user.Role = cs.Role
		}
		// passion
		if input.Passions != nil {
			user.Passions = input.Passions
		} else {
			user.Passions = cs.Passions
		}
		user.UpdatedAt = time.Now()
		r.Resolver.StudentStore[string(id)] = user
	} else {
		nid := uuid.New().String()
		user.ID = mypkg.UUID(nid)
		// role
		if input.Role != nil {
			user.Role = *input.Role
		} else {
			user.Role = model.UserTypeUser
		}
		// premium
		if input.IsGenius != nil {
			user.IsGenius = *input.IsGenius
		} else {
			user.IsGenius = false
		}
		// passions
		if input.Passions != nil {
			user.Passions = input.Passions
		} else {
			user.Passions = nil
		}
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		r.Resolver.StudentStore[nid] = user
	}
	res.Success = true
	res.Student = &user
	return &res, nil
}

func (r *queryResolver) Student(ctx context.Context, id mypkg.UUID) (*model.GetStudentResponse, error) {
	var res model.GetStudentResponse
	student, ok := r.Resolver.StudentStore[string(id)]

	if !ok {
		return nil, utils.ErrorResponse(ctx, "USER_NOT_FOUND", fmt.Errorf("user not found"))
	}
	res.Student = &student
	res.Success = true
	return &res, nil
}

func (r *queryResolver) Students(ctx context.Context, limit int) (*model.GetStudentsResponse, error) {
	var res model.GetStudentsResponse
	students := make([]*model.Student, 0)

	for i := range r.Resolver.StudentStore {
		if len(students) >= limit {
			break
		}
		u := r.Resolver.StudentStore[i]
		students = append(students, &u)
	}
	res.Students = students
	res.Success = true
	return &res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
