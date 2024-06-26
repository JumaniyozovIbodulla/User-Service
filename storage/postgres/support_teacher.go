package postgres

import (
	"context"
	sp "user/genproto/user_service/support_teachers"
	"user/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type supportTeacherRepo struct {
	db *pgxpool.Pool
}

func NewSupportTeacher(db *pgxpool.Pool) storage.SupportTeacherRepo {
	return &supportTeacherRepo{
		db: db,
	}
}

func (s *supportTeacherRepo) Create(ctx context.Context, req *sp.CreateSupportTeacher) (*sp.SupportTeacher, error) {
	panic("unimplemented")
}

func (s *supportTeacherRepo) Delete(ctx context.Context, req *sp.SupportTeacherPrimaryKey) (*sp.Empty, error) {
	panic("unimplemented")
}

func (s *supportTeacherRepo) GetAll(ctx context.Context, req *sp.GetListSupportTeachersRequest) (*sp.GetListSupportTeachersResponse, error) {
	panic("unimplemented")
}

func (s *supportTeacherRepo) GetById(ctx context.Context, req *sp.SupportTeacherPrimaryKey) (*sp.SupportTeacher, error) {
	panic("unimplemented")
}

func (s *supportTeacherRepo) Update(ctx context.Context, req *sp.UpdateSupportTeacher) (*sp.SupportTeacher, error) {
	panic("unimplemented")
}
