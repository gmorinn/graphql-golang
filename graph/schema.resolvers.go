package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-golang/graph/model"
	"graphql-golang/graph/mypkg"

	"github.com/google/uuid"
)

// UpdateStudent is the resolver for the UpdateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, input model.UpdateStudentInput) (*model.GetStudentResponse, error) {
	return r.Resolver.StudentService.UpdateStudent(ctx, &input)
}

// Signin is the resolver for the signin field.
func (r *mutationResolver) Signin(ctx context.Context, input model.SigninInput) (*model.JWTResponse, error) {
	return r.AuthService.Signin(ctx, &input)
}

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input model.SignupInput) (*model.JWTResponse, error) {
	return r.AuthService.Signup(ctx, &input)
}

// Refresh is the resolver for the refresh field.
func (r *mutationResolver) Refresh(ctx context.Context, refreshToken mypkg.JWT) (*model.JWTResponse, error) {
	return r.AuthService.RefreshToken(ctx, &refreshToken)
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, role model.UserType, id mypkg.UUID) (*model.GetStudentResponse, error) {
	return r.StudentService.UpdateRole(ctx, &role, &id)
}

// SingleUpload is the resolver for the singleUpload field.
func (r *mutationResolver) SingleUpload(ctx context.Context, file model.UploadInput) (*model.File, error) {
	return r.FileService.UploadSingleFile(ctx, &file)
}

// PostMessage is the resolver for the postMessage field.
func (r *mutationResolver) PostMessage(ctx context.Context, user string, content string) (string, error) {
	msg := &model.Message{
		ID:      uuid.New().String(),
		User:    user,
		Content: content,
	}
	r.ChatMessages = append(r.ChatMessages, msg)
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, msgs := range r.ChatObservers {
		msgs <- r.ChatMessages
	}
	return msg.ID, nil
}

// Student is the resolver for the student field.
func (r *queryResolver) Student(ctx context.Context, id mypkg.UUID) (*model.GetStudentResponse, error) {
	return r.StudentService.GetStudentByID(ctx, id)
}

// Students is the resolver for the students field.
func (r *queryResolver) Students(ctx context.Context, limit mypkg.NonNegativeInt, offset mypkg.NonNegativeInt) (*model.GetStudentsResponse, error) {
	return r.StudentService.GetStudents(ctx, limit, offset)
}

// Protected is the resolver for the protected field.
func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	return r.AuthService.Protected(ctx)
}

// Messages is the resolver for the messages field.
func (r *subscriptionResolver) Messages(ctx context.Context) (<-chan []*model.Message, error) {
	id := uuid.New().String()
	msgs := make(chan []*model.Message, 1)

	go func() {
		<-ctx.Done()
		r.mu.Lock()
		defer r.mu.Unlock()
		delete(r.ChatObservers, id)
		close(msgs)
	}()
	r.mu.Lock()
	r.ChatObservers[id] = msgs
	r.mu.Unlock()
	r.ChatObservers[id] <- r.ChatMessages

	return msgs, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
