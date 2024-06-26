package postgres

import (
	"context"
	ad "user/genproto/user_service/administrators"
	"user/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type adminRepo struct {
	db *pgxpool.Pool
}

func NewAdminRepo(db *pgxpool.Pool) storage.AdminRepo {
	return &adminRepo{
		db: db,
	}
}


func (a *adminRepo) Create(ctx context.Context, req *ad.CreateAdminstrator) (*ad.Adminstrator, error) {
	panic("unimplemented")
}

func (a *adminRepo) Delete(ctx context.Context, req *ad.AdminstratorPrimaryKey) (*ad.Empty, error) {
	panic("unimplemented")
}

func (a *adminRepo) GetAll(ctx context.Context, req *ad.GetListAdminstratorsRequest) (*ad.GetListAdminstratorsResponse, error) {
	panic("unimplemented")
}

func (a *adminRepo) GetById(ctx context.Context, req *ad.AdminstratorPrimaryKey) (*ad.Adminstrator, error) {
	panic("unimplemented")
}

func (a *adminRepo) Update(ctx context.Context, req *ad.UpdateAdminstrator) (*ad.Adminstrator, error) {
	panic("unimplemented")
}
