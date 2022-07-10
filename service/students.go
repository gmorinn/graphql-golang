package service

import (
	"context"
	config "graphql-golang/config"
	"graphql-golang/graph/model"
	"graphql-golang/graph/mypkg"
	db "graphql-golang/internal"
	"graphql-golang/utils"

	"github.com/google/uuid"
)

type IStudentService interface {
	CreateStudent(ctx context.Context, input *model.AddStudentInput) (*model.GetStudentResponse, error)
	UpdateStudent(ctx context.Context, input *model.UpdateStudentInput) (*model.GetStudentResponse, error)
	GetStudents(ctx context.Context, limit int, offset int) (*model.GetStudentsResponse, error)
	GetStudentByID(ctx context.Context, id mypkg.UUID) (*model.GetStudentResponse, error)
}

type StudentService struct {
	server *config.Server
}

func NewStudentService(server *config.Server) *StudentService {
	return &StudentService{
		server: server,
	}
}

func (s *StudentService) CreateStudent(ctx context.Context, input *model.AddStudentInput) (*model.GetStudentResponse, error) {
	var res model.Student

	var err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {

		stud, err := q.InsertStudent(ctx, db.InsertStudentParams{
			Email: string(input.Email),
			Name:  utils.NullS(input.Name),
		})

		if err != nil {
			return err
		}

		res = model.Student{
			ID:        mypkg.UUID(stud.ID.String()),
			Email:     mypkg.Email(stud.Email),
			Name:      stud.Name.String,
			CreatedAt: stud.CreatedAt,
			UpdatedAt: stud.UpdatedAt,
			DeletedAt: &stud.DeletedAt.Time,
			Role:      model.UserType(stud.Role),
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse(ctx, "TX_CREATE_STUDENT", err)
	}

	return &model.GetStudentResponse{
		Student: &res,
		Success: true,
	}, nil
}

func (s *StudentService) UpdateStudent(ctx context.Context, input *model.UpdateStudentInput) (*model.GetStudentResponse, error) {
	var res model.Student

	var err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		id := string(input.ID)

		if err := q.UpdateStudent(ctx, db.UpdateStudentParams{
			ID: uuid.MustParse(id),
		}); err != nil {
			return err
		}

		stud, err := q.GetStudentByID(ctx, uuid.MustParse(id))
		if err != nil {
			return err
		}

		res = model.Student{
			ID:        mypkg.UUID(stud.ID.String()),
			Email:     mypkg.Email(stud.Email),
			Name:      stud.Name.String,
			CreatedAt: stud.CreatedAt,
			UpdatedAt: stud.UpdatedAt,
			DeletedAt: &stud.DeletedAt.Time,
			Role:      model.UserType(stud.Role),
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse(ctx, "TX_UPDATE_STUDENT", err)
	}

	return &model.GetStudentResponse{
		Student: &res,
		Success: true,
	}, nil
}

func (s *StudentService) GetStudents(ctx context.Context, limit int, offset int) (*model.GetStudentsResponse, error) {
	var res []*model.Student = make([]*model.Student, 0)

	var err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		studs, err := q.Liststudents(ctx, db.ListstudentsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			return err
		}

		for _, stud := range studs {
			res = append(res, &model.Student{
				ID:        mypkg.UUID(stud.ID.String()),
				Email:     mypkg.Email(stud.Email),
				Name:      stud.Name.String,
				CreatedAt: stud.CreatedAt,
				UpdatedAt: stud.UpdatedAt,
				DeletedAt: &stud.DeletedAt.Time,
				Role:      model.UserType(stud.Role),
			})
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse(ctx, "TX_GET_STUDENTS", err)
	}

	return &model.GetStudentsResponse{
		Students: res,
		Success:  true,
	}, nil
}

func (s *StudentService) GetStudentByID(ctx context.Context, id mypkg.UUID) (*model.GetStudentResponse, error) {
	var res model.Student

	var err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		id := string(id)

		stud, err := q.GetStudentByID(ctx, uuid.MustParse(id))
		if err != nil {
			return err
		}

		res = model.Student{
			ID:        mypkg.UUID(stud.ID.String()),
			Email:     mypkg.Email(stud.Email),
			Name:      stud.Name.String,
			CreatedAt: stud.CreatedAt,
			UpdatedAt: stud.UpdatedAt,
			DeletedAt: &stud.DeletedAt.Time,
			Role:      model.UserType(stud.Role),
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse(ctx, "TX_GET_STUDENT_BY_ID", err)
	}

	return &model.GetStudentResponse{
		Student: &res,
		Success: true,
	}, nil
}
