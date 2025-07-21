package pgdb

import (
	"context"
	"fmt"
	"vk-server-task/internal/models"
	"vk-server-task/pkg/postgres"
)

const (
	adsTable        = "ads"
	paginationLimit = 30
)

type adsStorage struct {
	pg *postgres.Postgres
}

func NewAdsStorage(pg *postgres.Postgres) *adsStorage {
	return &adsStorage{
		pg: pg,
	}
}

func (s *adsStorage) Create(ctx context.Context, ad *models.Ads) (int, error) {
	var id int
	sql := fmt.Sprintf("INSERT INTO %s (user_id, title, description, image_url, price) VALUES ($1, $2, $3, $4, $5) RETURNING id", adsTable)
	err := s.pg.Pool.QueryRow(ctx, sql, ad.UserId, ad.Title, ad.Description, ad.ImageURL, ad.Price).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *adsStorage) Get(ctx context.Context, userId int, params models.AdsParams) ([]models.AdsResponse, error) {
	offset := paginationLimit * (params.Page - 1)

	cond := ""
	if params.MinPrice != nil || params.MaxPrice != nil {
		if params.MinPrice != nil && params.MaxPrice != nil {
			cond = fmt.Sprintf("WHERE price >= %d AND price <= %d", params.MinPrice, params.MaxPrice)
		} else if params.MinPrice != nil {
			cond = fmt.Sprintf("WHERE price >= %d", params.MinPrice)
		} else {
			cond = fmt.Sprintf("WHERE price <= %d", params.MaxPrice)
		}
	}

	sql := fmt.Sprintf("SELECT id, user_id, title, description, image_url, price, created_at FROM %s %s ORDER BY %s OFFSET $1 LIMIT $2", adsTable, cond, params.OrderField+" "+params.Order)

	rows, err := s.pg.Pool.Query(ctx, sql, offset, paginationLimit)
	if err != nil {
		return nil, fmt.Errorf("postgres.Get - r.pg.Pool.Query: %v", err)
	}
	defer rows.Close()

	var ads []models.AdsResponse
	for rows.Next() {
		var ad models.AdsResponse
		var id int
		err := rows.Scan(
			&ad.Id,
			&id,
			&ad.Title,
			&ad.Description,
			&ad.ImageURL,
			&ad.Price,
			&ad.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("postgres.Get - rows.Scan: %v", err)
		}

		if userId == id {
			ad.IsOwner = true
		}

		ads = append(ads, ad)
	}

	return ads, nil
}
