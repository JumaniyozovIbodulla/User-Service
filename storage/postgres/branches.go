package postgres

import (
	"context"
	br "user/genproto/user_service/branches"
	"user/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}


func NewBranchRepo(db *pgxpool.Pool) storage.BranchRepo {
	return &branchRepo{
		db: db,
	}
}

func (b *branchRepo) Create(ctx context.Context, req *br.CreateBranch) (*br.Branch, error) {
	panic("unimplemented")
}

func (b *branchRepo) Delete(ctx context.Context, req *br.BranchePrimaryKey) (*br.Empty, error) {
	panic("unimplemented")
}

func (b *branchRepo) GetAll(ctx context.Context, req *br.GetListBranchesRequest) (*br.GetListBranchesResponse, error) {
	panic("unimplemented")
}

func (b *branchRepo) GetById(ctx context.Context, req *br.BranchePrimaryKey) (*br.Branch, error) {
	panic("unimplemented")
}

func (b *branchRepo) Update(ctx context.Context, req *br.UpdateBranch) (*br.Branch, error) {
	panic("unimplemented")
}