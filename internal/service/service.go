package service

import (
	"context"
	"vk-server-task/internal/models"
	"vk-server-task/internal/storage"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Auth interface {
	ParseToken(accessToken string) (int, error)
	CreateUser(ctx context.Context, login, password string) (*models.User, string, error)
	LoginUser(ctx context.Context, login, password string) (*models.User, string, error)
}

type Ads interface {
	Create(ctx context.Context, userId int, params CreateRequest) (*models.Ads, error)
	Get(ctx context.Context, userId int, params models.AdsParams) ([]models.AdsResponse, error)
}

type Service struct {
	Auth
	Ads
}

func New(s *storage.Storage) *Service {
	return &Service{
		Auth: NewAuthService(s.Auth),
		Ads:  NewAdsService(s.Ads),
	}
}
