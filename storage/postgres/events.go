package postgres

import (
	"context"
	ev "user/genproto/user_service/events"
	"user/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type eventRepo struct {
	db *pgxpool.Pool
}

func NewEventRepo(db *pgxpool.Pool) storage.EventRepo {
	return &eventRepo{
		db: db,
	}
}

func (e *eventRepo) Create(ctx context.Context, req *ev.CreateEvent) (*ev.Event, error) {
	panic("unimplemented")
}

func (e *eventRepo) Delete(ctx context.Context, req *ev.EventPrimaryKey) (*ev.Empty, error) {
	panic("unimplemented")
}

func (e *eventRepo) GetAll(ctx context.Context, req *ev.GetListEventsRequest) (*ev.GetListEventsResponse, error) {
	panic("unimplemented")
}

func (e *eventRepo) GetById(ctx context.Context, req *ev.EventPrimaryKey) (*ev.Event, error) {
	panic("unimplemented")
}

func (e *eventRepo) Update(ctx context.Context, req *ev.UpdateEvent) (*ev.Event, error) {
	panic("unimplemented")
}

