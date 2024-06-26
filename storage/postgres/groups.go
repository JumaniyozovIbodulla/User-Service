package postgres

import (
	"context"
	gr "user/genproto/user_service/groups"
	"user/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type groupRepo struct {
	db *pgxpool.Pool
}

func NewGroupRepo(db *pgxpool.Pool) storage.GroupRepo {
	return &groupRepo{
		db: db,
	}
}


func (g *groupRepo) Create(ctx context.Context, req *gr.CreateGroup) (*gr.Group, error) {
	panic("unimplemented")
}

func (g *groupRepo) Delete(ctx context.Context, req *gr.GroupPrimaryKey) (*gr.Empty, error) {
	panic("unimplemented")
}

func (g *groupRepo) GetAll(ctx context.Context, req *gr.GetListGroupsRequest) (*gr.GetListGroupsResponse, error) {
	panic("unimplemented")
}

func (g *groupRepo) GetById(ctx context.Context, req *gr.GroupPrimaryKey) (*gr.Group, error) {
	panic("unimplemented")
}

func (g *groupRepo) Update(ctx context.Context, req *gr.UpdateGroup) (*gr.Group, error) {
	panic("unimplemented")
}