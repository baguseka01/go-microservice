package response

import (
	"time"

	"github.com/baguseka01/golang_microservice_hexagonal/business/film"
)

// GetFilmResponse Get film by ID response payload
type GetFilmResponse struct {
	ID         int       `json:"id"`
	UserID     int       `json:"film_id"`
	Title      string    `json:"title"`
	Genre      string    `json:"genre"`
	Artist     string    `json:"artist"`
	Rating     float32   `json:"rating"`
	ModifiedAt time.Time `json:"modified_at"`
	Version    int       `json:"version"`
}

// NewGetFilmResponse construct GetFilmResponse
func NewGetFilmResponse(film film.Film) *GetFilmResponse {
	var getFilmResponse GetFilmResponse

	getFilmResponse.ID = film.ID
	getFilmResponse.UserID = film.UserID
	getFilmResponse.Title = film.Title
	getFilmResponse.Genre = film.Genre
	getFilmResponse.Artist = film.Artist
	getFilmResponse.Rating = film.Rating
	getFilmResponse.ModifiedAt = film.ModifiedAt
	getFilmResponse.Version = film.Version

	return &getFilmResponse
}
