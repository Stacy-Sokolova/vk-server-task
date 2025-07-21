package storage

import (
	"context"
	"vk-server-task/internal/models"
	"vk-server-task/internal/storage/pgdb"
	"vk-server-task/pkg/postgres"
)

type Auth interface {
	GetUserByLogin(ctx context.Context, login string) (*models.User, error)
	CreateUser(ctx context.Context, login, password string) (int, error)
}

type Ads interface {
	Create(ctx context.Context, ad *models.Ads) (int, error)
	Get(ctx context.Context, userId int, params models.AdsParams) ([]models.AdsResponse, error)
}

type Storage struct {
	Auth
	Ads
}

func New(pg *postgres.Postgres) *Storage {
	return &Storage{
		Ads:  pgdb.NewAdsStorage(pg),
		Auth: pgdb.NewAuthStorage(pg),
	}
}
