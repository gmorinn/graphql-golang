package utils

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ErrorResponse(ctx context.Context, err string, error_code error) error {
	graphql.AddError(ctx, &gqlerror.Error{
		Extensions: map[string]interface{}{
			"err":        err,
			"error_code": error_code.Error(),
			"success":    false,
		},
	})
	return nil
}
