package postgresql

import (
	"context"
	"fmt"
	"net/url"

	"github.com/akmuhammetakmyradov/test/pkg/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgres(cfg *config.Configs) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Postgres.UserName, url.QueryEscape(cfg.Postgres.Password), cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DbName, cfg.Postgres.Sslmode)
	DbPool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Printf("err in connection postgres client: %v", err)
		return nil, err
	}

	if err = DbPool.Ping(context.Background()); err != nil {
		fmt.Printf("err in ping postgresql: %v", err)
		return nil, err
	}

	return DbPool, nil
}
