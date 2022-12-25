package response

import (
	"github.com/baguseka01/golang_microservice_hexagonal/api/paginator"
	"github.com/baguseka01/golang_microservice_hexagonal/business/film"
)

type getAllFilmResponse struct {
	Meta  paginator.Meta    `json:"meta"`
	Films []GetFilmResponse `json:"films"`
}

// NewGetAllFilmRespponse construct GetAllFilmResponse
func NewGetAllFilmRespponse(films []film.Film, page int, rowPerPage int) getAllFilmResponse {

	var (
		lenFilms = len(films)
	)

	getAllFilmResponse := getAllFilmResponse{}
	getAllFilmResponse.Meta.BuildMeta(lenFilms, page, rowPerPage)

	for index, value := range films {
		if index == getAllFilmResponse.Meta.RowPerPage {
			continue
		}
		var getFilmResponse GetFilmResponse

		getFilmResponse.ID = value.ID
		getFilmResponse.UserID = value.UserID
		getFilmResponse.Title = value.Title
		getFilmResponse.Genre = value.Genre
		getFilmResponse.Artist = value.Artist
		getFilmResponse.Rating = value.Rating
		getFilmResponse.ModifiedAt = value.ModifiedAt
		getFilmResponse.Version = value.Version

		getAllFilmResponse.Films = append(getAllFilmResponse.Films, getFilmResponse)
	}

	if getAllFilmResponse.Films == nil {
		getAllFilmResponse.Films = []GetFilmResponse{}
	}

	return getAllFilmResponse

}
