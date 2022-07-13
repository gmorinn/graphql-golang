package graph

import (
	"graphql-golang/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	StudentService service.IStudentService
	AuthService    service.IAuthService
	FileService    service.IFileService
}
