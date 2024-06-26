package postgres

import (
	"context"
	br "user/genproto/user_service/branches"
	"user/storage"

	"github.com/google/uuid"
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
	id := uuid.New()

	_, err := b.db.Exec(ctx, `
		INSERT INTO
			branches(id, name, location, super_admin_id)
		VALUES(
		$1, 
		$2, 
		ST_SetSRID(ST_MakePoint($3, $4), 4326),
		$5);`,
		id,
		req.Name,
		req.Location.Latitude,
		req.Location.Longitude,
		req.SuperAdminId)

	if err != nil {
		return nil, err
	}

	branch, err := b.GetById(ctx, &br.BranchePrimaryKey{Id: id.String()})

	if err != nil {
		return nil, err
	}

	return branch, nil
}

func (b *branchRepo) Delete(ctx context.Context, req *br.BranchePrimaryKey) (*br.Empty, error) {
	_, err := b.db.Exec(ctx, `DELETE FROM branches WHERE id = $1;`, req.Id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (b *branchRepo) GetAll(ctx context.Context, req *br.GetListBranchesRequest) (*br.GetListBranchesResponse, error) {
	resp := &br.GetListBranchesResponse{}
	filter := ""

	if req.Search != "" {
		filter = ` AND name ILIKE '%` + req.Search + `%' `
	}

	rows, err := b.db.Query(ctx, `
	SELECT 
		id,
		name,
		ST_Y(location) AS latitude, 
		ST_X(location) AS longitude,
		super_admin_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM
		branches
	WHERE
		TRUE `+filter+`
	OFFSET
		$1
	LIMIT
		$2;`, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		branch := br.Branch{}

		err := rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Location.Latitude,
			&branch.Location.Longitude,
			&branch.SuperAdminId,
			&branch.IsActive,
			&branch.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Branches = append(resp.Branches, &branch)
	}

	defer rows.Close()

	err = b.db.QueryRow(ctx, `SELECT COUNT(*) FROM branches WHERE TRUE `+filter+`; `).Scan(&resp.Count)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *branchRepo) GetById(ctx context.Context, req *br.BranchePrimaryKey) (*br.Branch, error) {
	row := b.db.QueryRow(ctx, `
	SELECT
		id,
		name,
		ST_Y(location) AS latitude, 
		ST_X(location) AS longitude,
		super_admin_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at
	FROM
		branches
	WHERE
		id = $1;`, req.Id)
	
	var (
		branch br.Branch
		lat, long float64
	)

	err := row.Scan(
		&branch.Id,
		&branch.Name,
		&lat,
		&long,
		&branch.SuperAdminId,
		&branch.IsActive,
		&branch.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	branch.Location.Latitude = lat
	branch.Location.Longitude = long

	return &branch, nil
}

func (b *branchRepo) Update(ctx context.Context, req *br.UpdateBranch) (*br.Branch, error) {
	branch := &br.Branch{}

	row := b.db.QueryRow(ctx, `
	UPDATE 
		branches
	SET
		id = $2,
		name = $3,
		location = ST_SetSRID(ST_MakePoint($4, $5)
		super_admin_id = $6,
		is_active = $7
	WHERE
		id = $1
	RETURNING
		id,
		name,
		ST_Y(location) AS latitude, 
		ST_X(location) AS longitude,
		super_admin_id,
		is_active,
		TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS TZH:TZM') AS created_at;`,
		req.Id,
		req.Name,
		req.Location.Latitude,
		req.Location.Longitude,
		req.SuperAdminId,
		req.IsActive)

	err := row.Scan(
		&branch.Id,
		&branch.Name,
		&branch.Location.Latitude,
		&branch.Location.Longitude,
		&branch.SuperAdminId,
		&branch.IsActive,
		&branch.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return branch, nil
}
