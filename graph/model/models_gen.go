// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"graphql-golang/graph/mypkg"
	"io"
	"strconv"
	"time"
)

type Response interface {
	IsResponse()
}

// Interface that have the mandatory fields of the user in all projects
type User interface {
	IsUser()
}

type Competence struct {
	Name string `json:"name"`
}

type CompetenceInput struct {
	Name string `json:"name"`
}

// if there is an error, return it or null
type ErrorResponse struct {
	Err       string `json:"err"`
	ErrorCode string `json:"error_code"`
}

// Response when you get a student
type GetStudentResponse struct {
	// if the request was successful or not, return always a value
	Success bool `json:"success"`
	// return the student if the request was successful
	Student *Student `json:"student"`
}

func (GetStudentResponse) IsResponse() {}

// Response when you get many students
type GetStudentsResponse struct {
	// if the request was successful or not, return always a value
	Success bool `json:"success"`
	// return an array of student if the request was successful or null if there is an error or no students
	Students []*Student `json:"students"`
}

func (GetStudentsResponse) IsResponse() {}

// All fields that represent a student
type Student struct {
	Name        string        `json:"name"`
	Email       mypkg.Email   `json:"email"`
	ID          mypkg.UUID    `json:"id"`
	Age         int           `json:"age"`
	URL         mypkg.URL     `json:"url"`
	Gpa         float64       `json:"gpa"`
	Passions    []string      `json:"passions"`
	IsGenius    bool          `json:"is_genius"`
	Role        UserType      `json:"role"`
	CreatedAt   time.Time     `json:"created_at"`
	DeletedAt   time.Time     `json:"deleted_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Competences []*Competence `json:"competences"`
}

func (Student) IsUser() {}

type StudentInput struct {
	Name     string      `json:"name"`
	ID       *mypkg.UUID `json:"id"`
	Age      int         `json:"age"`
	Email    mypkg.Email `json:"email"`
	URL      mypkg.URL   `json:"url"`
	Gpa      float64     `json:"gpa"`
	Passions []string    `json:"passions"`
	IsGenius *bool       `json:"is_genius"`
	Role     *UserType   `json:"role"`
}

type UserType string

const (
	// User can have access to all data
	UserTypeAdmin UserType = "ADMIN"
	// User can access specific data but not all
	UserTypePro UserType = "PRO"
	// User can only see their own data
	UserTypeUser UserType = "USER"
)

var AllUserType = []UserType{
	UserTypeAdmin,
	UserTypePro,
	UserTypeUser,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeAdmin, UserTypePro, UserTypeUser:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
