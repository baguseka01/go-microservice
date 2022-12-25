package request

import "github.com/baguseka01/golang_microservice_hexagonal/business/film"

// InsertFilmRequest create Film request payload
type InsertFilmRequest struct {
	Title  string  `json:"title"`
	Genre  string  `json:"genre"`
	Artist string  `json:"artist"`
	Rating float32 `json:"rating"`
}

// ToUpsertFilmSpec convert into film,UpsertFilmSpec object
func (req *InsertFilmRequest) ToUpsertFilmSpec(userID int) *film.InsertFilmSpec {
	var insertFilmSpec film.InsertFilmSpec

	insertFilmSpec.UserID = userID
	insertFilmSpec.Title = req.Title
	insertFilmSpec.Genre = req.Genre
	insertFilmSpec.Artist = req.Artist
	insertFilmSpec.Rating = req.Rating

	return &insertFilmSpec
}
