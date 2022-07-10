package service

import (
	"context"
	"fmt"
	"graphql-golang/graph/model"
	"graphql-golang/graph/mypkg"
	"graphql-golang/utils"
	"time"

	"github.com/google/uuid"
)

type IStudentService interface {
	CreateStudent(ctx context.Context, input *model.StudentInput) (*model.GetStudentResponse, error)
	UpdateStudent(ctx context.Context, input *model.StudentInput) (*model.GetStudentResponse, error)
	GetStudents(ctx context.Context, limit int) (*model.GetStudentsResponse, error)
	GetStudentByID(ctx context.Context, id mypkg.UUID) (*model.GetStudentResponse, error)
}

type StudentService struct {
	StudentStore map[string]*model.Student
}

func NewStudentService(store map[string]*model.Student) *StudentService {
	return &StudentService{
		StudentStore: store,
	}
}

func (s *StudentService) CreateStudent(ctx context.Context, input *model.StudentInput) (*model.GetStudentResponse, error) {
	var user *model.Student
	nid := uuid.New().String()
	user.ID = mypkg.UUID(nid)
	user.Email = input.Email
	user.Name = input.Name
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.DeletedAt = nil
	s.StudentStore[nid] = user
	return &model.GetStudentResponse{
		Student: user,
		Success: true,
	}, nil
}

func (s *StudentService) UpdateStudent(ctx context.Context, input *model.StudentInput) (*model.GetStudentResponse, error) {
	var user *model.Student
	id := *input.ID
	cs, ok := s.StudentStore[string(id)]
	if !ok {
		return nil, utils.ErrorResponse(ctx, "USER_NOT_FOUND", fmt.Errorf("user not found"))
	}
	if input.Name != nil {
		user.Name = input.Name
	} else {
		user.Name = cs.Name
	}
	user.UpdatedAt = time.Now()
	s.StudentStore[string(id)] = user
	return &model.GetStudentResponse{
		Student: user,
		Success: true,
	}, nil
}

func (s *StudentService) GetStudents(ctx context.Context, limit int) (*model.GetStudentsResponse, error) {
	var res model.GetStudentsResponse
	students := make([]*model.Student, 0)

	for i := range s.StudentStore {
		if len(students) >= limit {
			break
		}
		u := s.StudentStore[i]
		students = append(students, u)
	}
	res.Students = students
	res.Success = true
	return &res, nil
}

func (s *StudentService) GetStudentByID(ctx context.Context, id mypkg.UUID) (*model.GetStudentResponse, error) {
	var res model.GetStudentResponse
	student, ok := s.StudentStore[string(id)]
	if !ok {
		return nil, utils.ErrorResponse(ctx, "USER_NOT_FOUND", fmt.Errorf("user not found"))
	}
	res.Student = student
	res.Success = true
	return &res, nil
}
