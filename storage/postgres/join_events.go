package postgres

import (
	"context"
	ej "user/genproto/user_service/join_events"
	"user/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type joinEventRepo struct {
	db *pgxpool.Pool
}


func NewEventJoinRepo(db *pgxpool.Pool) storage.EventJoinRepo {
	return &joinEventRepo{
		db: db,
	}
}


func (j *joinEventRepo) Create(ctx context.Context, req *ej.CreateJoinEvent) (*ej.JoinEvent, error) {
	panic("unimplemented")
}

func (j *joinEventRepo) GetById(ctx context.Context, req *ej.JoinEventPrimaryKey) (*ej.JoinEvent, error) {
	panic("unimplemented")
}