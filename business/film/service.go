package film

import (
	"time"

	"github.com/baguseka01/golang_microservice_hexagonal/business"
	"github.com/baguseka01/golang_microservice_hexagonal/util/validator"
)

// InsertFilmSpec create film spec
type InsertFilmSpec struct {
	UserID int     `validate:"required"`
	Title  string  `validate:"required"`
	Genre  string  `validate:"required"`
	Artist string  `validate:"required"`
	Rating float32 `validate:"required"`
}

// <====================================== THE IMPLEMENTATION OF THOSE INTERFACE PUT BELOW ======================================>

type service struct {
	repository Repository
}

// NewService Construct film service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

// FindFilmByID Get film by given ID, return nil if not exist
func (s *service) FindFilmByID(id int, userID int) (*Film, error) {
	return s.repository.FindFilmByID(id, userID)
}

// FindAllFilm Get All Film, will be return empty array if no data or error accurent
func (s *service) FindAllFilm(userID int, skip int, rowPerPage int) ([]Film, error) {

	film, err := s.repository.FindAllFilm(userID, skip, rowPerPage)
	if err != nil {
		return []Film{}, err
	}

	return film, err
}

// Insert Create new film and store into database
func (s *service) InsertFilm(insertFilmSpec InsertFilmSpec, createBy string) error {

	err := validator.GetValidator().Struct(insertFilmSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	film := NewFilm(
		0,
		insertFilmSpec.UserID,
		insertFilmSpec.Title,
		insertFilmSpec.Genre,
		insertFilmSpec.Artist,
		insertFilmSpec.Rating,
		createBy,
		time.Now(),
	)

	err = s.repository.InsertFilm(film)
	if err != nil {
		return err
	}

	return nil
}

// UpdateFilm will update found film, if not exist will be return error
func (s *service) UpdateFilm(id int, userID int, title string, modifiedBy string, currentVersion int) error {

	film, err := s.repository.FindFilmByID(id, userID)
	if err != nil {
		return err
	} else if film == nil {
		return business.ErrNotFound
	} else if film.Version != currentVersion {
		return business.ErrHasBeenModified
	}

	modifiedFilm := film.ModifyFilm(title, time.Now(), modifiedBy)

	return s.repository.UpdateFilm(modifiedFilm, currentVersion)
}
