package models

import "time"

type Ads struct {
	Id          int       `json:"id" db:"id"`
	UserId      int       `json:"-" db:"user_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description" binding:"required"`
	ImageURL    string    `json:"image_url" db:"image_url" binding:"required"`
	Price       float64   `json:"price" db:"price" binding:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type AdsParams struct {
	Page       int
	OrderField string
	Order      string
	MinPrice   *float64
	MaxPrice   *float64
}

type AdsResponse struct {
	Id          int
	IsOwner     bool
	Title       string
	Description string
	ImageURL    string
	Price       float64
	CreatedAt   time.Time
}
