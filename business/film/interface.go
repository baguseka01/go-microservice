package film

//  Service outgoing port for film
type Service interface {
	// FindFilmByID if data not found will return nil without error
	FindFilmByID(id int, userID int) (*Film, error)

	// FindAllFilm find all film with given specific pace and row per page, will return empty slice instead of nil
	FindAllFilm(userID int, skip int, rowPerPage int) ([]Film, error)

	// InsertFilm Insert new Film into storage
	InsertFilm(insertFilmSpec InsertFilmSpec, createdBy string) error

	// UpdateFilm if data not found will return error
	UpdateFilm(id int, userID int, title string, modifiiedBy string, currentVersion int) error
}

type Repository interface {
	// FindFilmByID if data not found will return nil without error
	FindFilmByID(id int, userID int) (*Film, error)

	// FindAllFilm find all film with given specific pace and row per page, will return empty slice instead of nil
	FindAllFilm(userID int, skip int, rowPerPage int) ([]Film, error)

	// InsertFilm Insert new Film into storage
	InsertFilm(film Film) error

	// UpdateFilm if data not found will return error
	UpdateFilm(film Film, currentVersion int) error
}