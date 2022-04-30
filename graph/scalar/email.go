package scalar

import (
	"fmt"
	"io"
	"regexp"
)

type Email string

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (value *Email) UnmarshalGQL(v interface{}) error {
	data, ok := v.(string)
	if !ok {
		return fmt.Errorf("Email must be a string")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if emailRegex.MatchString(data) {
		*value = Email(data)
	} else {
		return fmt.Errorf("Email must be a valid email")
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (value Email) MarshalGQL(w io.Writer) {
	w.Write([]byte(value))
}
