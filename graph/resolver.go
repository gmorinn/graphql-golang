package graph

import (
	"graphql-golang/graph/model"
	"graphql-golang/service"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	StudentService service.IStudentService
	AuthService    service.IAuthService
	FileService    service.IFileService

	// All messages since launching the GraphQL endpoint
	ChatMessages []*model.Message
	// All active subscriptions
	ChatObservers map[string]chan []*model.Message
	mu            sync.Mutex
}
