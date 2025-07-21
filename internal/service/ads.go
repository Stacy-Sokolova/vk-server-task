package service

import (
	"context"
	"vk-server-task/internal/models"
	"vk-server-task/internal/storage"

	"github.com/sirupsen/logrus"
)

type CreateRequest struct {
	Title       string  `json:"title" binding:"required,min=1,max=100"`
	Description string  `json:"description" binding:"required,min=1,max=1000"`
	ImageURL    string  `json:"image_url" binding:"required,url,max=2048,imageurl"`
	Price       float64 `json:"price" binding:"required,gt=0"`
}

type AdsService struct {
	storage storage.Ads
}

func NewAdsService(s storage.Ads) *AdsService {
	return &AdsService{storage: s}
}

func (s *AdsService) Create(ctx context.Context, userId int, params *CreateRequest) (*models.Ads, error) {
	ad := &models.Ads{
		Title:       params.Title,
		Description: params.Description,
		ImageURL:    params.ImageURL,
		Price:       params.Price,
		UserId:      userId,
	}
	id, err := s.storage.Create(ctx, ad)
	if err != nil {
		return nil, err
	}

	ad.Id = id

	logrus.Infof("new ad %d successfully created", id)
	return ad, nil
}

func (s *AdsService) Get(ctx context.Context, userId int, params models.AdsParams) ([]models.AdsResponse, error) {
	return s.storage.Get(ctx, userId, params)
}
