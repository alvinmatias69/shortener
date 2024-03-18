package repository

import (
	"context"
	"errors"

	"github.com/alvinmatias69/shortener/internal/constants"
	"github.com/alvinmatias69/shortener/internal/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) Get(ctx context.Context, hash string) (entity.UrlPayload, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, hash, long_url FROM mst_url_shortener WHERE hash = $1", hash)
	if err != nil {
		return entity.UrlPayload{}, err
	}

	payload, err := pgx.CollectOneRow[entity.UrlPayload](rows, pgx.RowToStructByName[entity.UrlPayload])
	if errors.Is(err, pgx.ErrNoRows) {
		return entity.UrlPayload{}, constants.UrlNotFound
	}

	if err != nil {
		return entity.UrlPayload{}, err
	}

	return payload, nil
}

func (r *Repository) GetByLongUrl(ctx context.Context, longUrl string) (entity.UrlPayload, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, hash, long_url FROM mst_url_shortener WHERE long_url = $1", longUrl)
	if err != nil {
		return entity.UrlPayload{}, err
	}

	payload, err := pgx.CollectOneRow[entity.UrlPayload](rows, pgx.RowToStructByName[entity.UrlPayload])
	if errors.Is(err, pgx.ErrNoRows) {
		return entity.UrlPayload{}, constants.UrlNotFound
	}

	if err != nil {
		return entity.UrlPayload{}, err
	}

	return payload, nil
}

func (r *Repository) Create(ctx context.Context, payload entity.UrlPayload) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "INSERT INTO mst_url_shortener(id, hash, long_url) VALUES ($1, $2, $3)",
		payload.Id, payload.Hash, payload.LongUrl)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
