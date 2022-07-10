package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-golang/graph/model"
	"graphql-golang/graph/mypkg"
)

// AddStudent is the resolver for the addStudent field.
func (r *mutationResolver) AddStudent(ctx context.Context, input model.AddStudentInput) (*model.GetStudentResponse, error) {
	return r.Resolver.StudentService.CreateStudent(ctx, &input)
}

// UpdateStudent is the resolver for the UpdateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, input model.UpdateStudentInput) (*model.GetStudentResponse, error) {
	return r.Resolver.StudentService.UpdateStudent(ctx, &input)
}

// Student is the resolver for the student field.
func (r *queryResolver) Student(ctx context.Context, id mypkg.UUID) (*model.GetStudentResponse, error) {
	return r.StudentService.GetStudentByID(ctx, id)
}

// Students is the resolver for the students field.
func (r *queryResolver) Students(ctx context.Context, limit int, offset int) (*model.GetStudentsResponse, error) {
	return r.StudentService.GetStudents(ctx, limit, offset)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
