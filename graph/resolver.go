package graph

import "graphql-golang/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	StudentStore map[string]model.Student
}