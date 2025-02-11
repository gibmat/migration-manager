package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/FuturFusion/migration-manager/internal/migration/repo"
)

type target struct {
	db repo.DBTX
}

var _ migration.TargetRepo = &target{}

func NewTarget(db repo.DBTX) *target {
	return &target{
		db: db,
	}
}

func (t target) Create(ctx context.Context, in migration.Target) (migration.Target, error) {
	const sqlInsert = `
INSERT INTO targets (name, endpoint, tls_client_key, tls_client_cert, oidc_tokens, insecure)
VALUES(:name, :endpoint, :tls_client_key, :tls_client_cert, :oidc_tokens, :insecure)
RETURNING id, name, endpoint, tls_client_key, tls_client_cert, oidc_tokens, insecure;
`

	marshalledOIDCTokens, err := json.Marshal(in.OIDCTokens)
	if err != nil {
		return migration.Target{}, err
	}

	row := t.db.QueryRowContext(ctx, sqlInsert,
		sql.Named("name", in.Name),
		sql.Named("endpoint", in.Endpoint),
		sql.Named("tls_client_key", in.TLSClientKey),
		sql.Named("tls_client_cert", in.TLSClientCert),
		sql.Named("oidc_tokens", marshalledOIDCTokens),
		sql.Named("insecure", in.Insecure),
	)
	if row.Err() != nil {
		return migration.Target{}, mapErr(row.Err())
	}

	return scanTarget(row)
}

func (t target) GetAll(ctx context.Context) (migration.Targets, error) {
	const sqlGetAll = `SELECT id, name, endpoint, tls_client_key, tls_client_cert, oidc_tokens, insecure FROM targets ORDER BY name;`

	rows, err := t.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return nil, mapErr(err)
	}

	defer func() { _ = rows.Close() }()

	var targets migration.Targets
	for rows.Next() {
		target, err := scanTarget(rows)
		if err != nil {
			return nil, err
		}

		targets = append(targets, target)
	}

	if rows.Err() != nil {
		return nil, mapErr(rows.Err())
	}

	return targets, nil
}

func (t target) GetAllNames(ctx context.Context) ([]string, error) {
	const sqlGetAllNames = `SELECT name FROM targets ORDER BY name`

	rows, err := t.db.QueryContext(ctx, sqlGetAllNames)
	if err != nil {
		return nil, mapErr(err)
	}

	defer func() { _ = rows.Close() }()

	var targetNames []string
	for rows.Next() {
		var targetName string
		err := rows.Scan(&targetName)
		if err != nil {
			return nil, mapErr(err)
		}

		targetNames = append(targetNames, targetName)
	}

	if rows.Err() != nil {
		return nil, mapErr(rows.Err())
	}

	return targetNames, nil
}

func (t target) GetByID(ctx context.Context, id int) (migration.Target, error) {
	const sqlGetByID = `SELECT id, name, endpoint, tls_client_key, tls_client_cert, oidc_tokens, insecure FROM targets WHERE id=:id;`

	row := t.db.QueryRowContext(ctx, sqlGetByID, sql.Named("id", id))
	if row.Err() != nil {
		return migration.Target{}, mapErr(row.Err())
	}

	return scanTarget(row)
}

func (t target) GetByName(ctx context.Context, name string) (migration.Target, error) {
	const sqlGetByName = `SELECT id, name, endpoint, tls_client_key, tls_client_cert, oidc_tokens, insecure FROM targets WHERE name=:name;`

	row := t.db.QueryRowContext(ctx, sqlGetByName, sql.Named("name", name))
	if row.Err() != nil {
		return migration.Target{}, mapErr(row.Err())
	}

	return scanTarget(row)
}

func (t target) UpdateByID(ctx context.Context, in migration.Target) (migration.Target, error) {
	const sqlUpdate = `
UPDATE targets SET name=:name, endpoint=:endpoint, tls_client_key=:tls_client_key, tls_client_cert=:tls_client_cert, oidc_tokens=:oidc_tokens, insecure=:insecure
WHERE id=:id
RETURNING id, name, endpoint, tls_client_key, tls_client_cert, oidc_tokens, insecure;
`

	marshalledOIDCTokens, err := json.Marshal(in.OIDCTokens)
	if err != nil {
		return migration.Target{}, err
	}

	row := t.db.QueryRowContext(ctx, sqlUpdate,
		sql.Named("name", in.Name),
		sql.Named("endpoint", in.Endpoint),
		sql.Named("tls_client_key", in.TLSClientKey),
		sql.Named("tls_client_cert", in.TLSClientCert),
		sql.Named("oidc_tokens", marshalledOIDCTokens),
		sql.Named("insecure", in.Insecure),
		sql.Named("id", in.ID),
	)
	if row.Err() != nil {
		return migration.Target{}, mapErr(row.Err())
	}

	return scanTarget(row)
}

func scanTarget(row interface{ Scan(dest ...any) error }) (migration.Target, error) {
	var target migration.Target
	var marshalledOIDCTokens []byte
	err := row.Scan(
		&target.ID,
		&target.Name,
		&target.Endpoint,
		&target.TLSClientKey,
		&target.TLSClientCert,
		&marshalledOIDCTokens,
		&target.Insecure,
	)
	if err != nil {
		return migration.Target{}, mapErr(err)
	}

	err = json.Unmarshal(marshalledOIDCTokens, &target.OIDCTokens)
	if err != nil {
		return migration.Target{}, err
	}

	return target, nil
}

func (t target) DeleteByName(ctx context.Context, name string) error {
	const sqlDelete = `DELETE FROM targets WHERE name=:name;`

	result, err := t.db.ExecContext(ctx, sqlDelete, sql.Named("name", name))
	if err != nil {
		return mapErr(err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return mapErr(err)
	}

	if affectedRows == 0 {
		return migration.ErrNotFound
	}

	return nil
}
