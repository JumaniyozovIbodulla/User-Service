package postgres

import (
	"context"
	mn "user/genproto/user_service/managers"
	"user/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type managerRepo struct {
	db *pgxpool.Pool
}

func NewManagerRepo(db *pgxpool.Pool) storage.ManagerRepo {
	return &managerRepo{
		db: db,
	}
}

func (m *managerRepo) Create(ctx context.Context, req *mn.CreateManager) (*mn.Manager, error) {
	panic("unimplemented")
}

func (m *managerRepo) Delete(ctx context.Context, req *mn.ManagerPrimaryKey) (*mn.Empty, error) {
	panic("unimplemented")
}

func (m *managerRepo) GetAll(ctx context.Context, req *mn.GetListManagersRequest) (*mn.GetListManagersResponse, error) {
	panic("unimplemented")
}

func (m *managerRepo) GetById(ctx context.Context, req *mn.ManagerPrimaryKey) (*mn.Manager, error) {
	panic("unimplemented")
}

func (m *managerRepo) Update(ctx context.Context, req *mn.UpdateManager) (*mn.Manager, error) {
	panic("unimplemented")
}

