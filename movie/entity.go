package movie

import (
	"time"
)

type Movie struct{
	ID int `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Rating float32 `json:"rating" validate:"required"`
	Image string `json:"image" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}