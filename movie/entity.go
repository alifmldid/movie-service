package movie

import "time"

type Movie struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Rating float32 `json:"rating"`
	Image string `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}