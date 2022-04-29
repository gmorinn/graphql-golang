package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-golang/graph/generated"
	"graphql-golang/graph/model"
	"time"

	"github.com/google/uuid"
)

func (r *mutationResolver) AddOrUpdateStudent(ctx context.Context, input model.StudentInput) (*model.Student, error) {
	id := input.ID
	var user model.Student
	user.Name = input.Name
	user.Age = input.Age
	user.Gpa = input.Gpa

	n := len(r.Resolver.StudentStore)
	if n == 0 {
		r.Resolver.StudentStore = make(map[string]model.Student)
	}

	if id != nil {
		cs, ok := r.Resolver.StudentStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		// is_premium
		if input.IsPremium != nil {
			user.IsPremium = *input.IsPremium
		} else {
			user.IsPremium = cs.IsPremium
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
		r.Resolver.StudentStore[*id] = user
	} else {
		nid := uuid.New().String()
		user.ID = nid
		// role
		if input.Role != nil {
			user.Role = *input.Role
		} else {
			user.Role = model.UserTypeUser
		}
		// premium
		if input.IsPremium != nil {
			user.IsPremium = *input.IsPremium
		} else {
			user.IsPremium = false
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

	return &user, nil
}

func (r *queryResolver) Student(ctx context.Context, id string) (*model.Student, error) {
	res, ok := r.Resolver.StudentStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &res, nil
}

func (r *queryResolver) Students(ctx context.Context, limit int) ([]*model.Student, error) {
	var res []*model.Student = make([]*model.Student, 0)
	var u model.Student

	for i := range r.Resolver.StudentStore {
		if len(res) >= limit {
			break
		}
		u = r.Resolver.StudentStore[i]
		res = append(res, &u)

	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
