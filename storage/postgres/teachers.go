package postgres

import (
	"context"
	tr "user/genproto/user_service/teachers"
	"user/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type teacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacherRepo(db *pgxpool.Pool) storage.TeacherRepo {
	return &teacherRepo{
		db: db,
	}
}


func (t *teacherRepo) Create(ctx context.Context, req *tr.CreateTeacher) (*tr.Teacher, error) {
	panic("unimplemented")
}

func (t *teacherRepo) Delete(ctx context.Context, req *tr.TeacherPrimaryKey) (*tr.Empty, error) {
	panic("unimplemented")
}

func (t *teacherRepo) GetAll(ctx context.Context, req *tr.GetListTeachersRequest) (*tr.GetListTeachersResponse, error) {
	panic("unimplemented")
}

func (t *teacherRepo) GetById(ctx context.Context, req *tr.TeacherPrimaryKey) (*tr.Teacher, error) {
	panic("unimplemented")
}

func (t *teacherRepo) Update(ctx context.Context, req *tr.UpdateTeacher) (*tr.Teacher, error) {
	panic("unimplemented")
}