package postgres

import (
	"context"
	st "user/genproto/user_service/students"
	"user/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type studentRepo struct {
	db *pgxpool.Pool
}

func NewStudentRepo(db *pgxpool.Pool) storage.StudentRepo {
	return &studentRepo{
		db: db,
	}
}

func (s *studentRepo) Create(ctx context.Context, req *st.CreateStudent) (*st.Student, error) {
	panic("unimplemented")
}

func (s *studentRepo) Delete(ctx context.Context, req *st.StudentPrimaryKey) (*st.Empty, error) {
	panic("unimplemented")
}

func (s *studentRepo) GetAll(ctx context.Context, req *st.GetListStudentsRequest) (*st.GetListStudentsResponse, error) {
	panic("unimplemented")
}

func (s *studentRepo) GetById(ctx context.Context, req *st.StudentPrimaryKey) (*st.Student, error) {
	panic("unimplemented")
}

func (s *studentRepo) Update(ctx context.Context, req *st.UpdateStudent) (*st.Student, error) {
	panic("unimplemented")
}
