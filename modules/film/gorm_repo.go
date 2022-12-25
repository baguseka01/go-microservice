package film

import (
	"time"

	"github.com/baguseka01/golang_microservice_hexagonal/business/film"
	"github.com/baguseka01/golang_microservice_hexagonal/modules/user"
	"gorm.io/gorm"
)

// GormRepository The implentation of film.Repository object
type GormRepository struct {
	DB *gorm.DB
}

type FilmTable struct {
	ID         int            `gorm:"id"`
	UserID     int            `gorm:"user_id"`
	Title      string         `gorm:"title"`
	Genre      string         `gorm:"genre"`
	Artist     string         `gorm:"artist"`
	Rating     float32        `gorm:"rating"`
	CreatedAt  time.Time      `gorm:"created_at"`
	CreatedBy  string         `gorm:"created_by"`
	ModifiedAt time.Time      `gorm:"modified_at"`
	ModifiedBy string         `gorm:"modified_by"`
	Version    int            `gorm:"version"`
	User       user.UserTable `gorm:"foreignKey:UserID"`
}

func newFilmTable(film film.Film) *FilmTable {

	return &FilmTable{
		film.ID,
		film.UserID,
		film.Title,
		film.Genre,
		film.Artist,
		film.Rating,
		film.CreatedAt,
		film.CreatedBy,
		film.ModifiedAt,
		film.ModifiedBy,
		film.Version,
		user.UserTable{},
	}
}

func (col *FilmTable) ToFilm() film.Film {
	var film film.Film

	film.ID = col.ID
	film.UserID = col.UserID
	film.Title = col.Title
	film.Genre = col.Genre
	film.Artist = col.Artist
	film.Rating = col.Rating
	film.CreatedAt = col.CreatedAt
	film.CreatedBy = col.CreatedBy
	film.ModifiedAt = col.ModifiedAt
	film.ModifiedBy = col.ModifiedBy
	film.Version = col.Version

	return film
}

// NewGormDBRepository Generate Gorm DB film repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

// FindFilmByID if data not found will return nil without error
func (repo *GormRepository) FindFilmByID(id int, userID int) (*film.Film, error) {
	var filmData FilmTable

	err := repo.DB.Where("id = ?", id).Where("user_id = ?", userID).First(&filmData).Error
	if err != nil {
		return nil, err
	}

	film := filmData.ToFilm()

	return &film, nil
}

// FindAllFilm find all film with given specific page and row, will return empty slice instead of nil
func (repo *GormRepository) FindAllFilm(userID int, skip int, rowPerPage int) ([]film.Film, error) {
	var films []FilmTable

	err := repo.DB.Where("user_id = ?", userID).Offset(skip).Limit(rowPerPage).Find(&films).Error
	if err != nil {
		return nil, err
	}

	var result []film.Film
	for _, value := range films {
		result = append(result, value.ToFilm())
	}

	return result, nil

}

// InsertFilm Insert new film into storage
func (repo *GormRepository) InsertFilm(film film.Film) error {

	filmData := newFilmTable(film)
	filmData.ID = 0

	err := repo.DB.Create(filmData).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateFilm Update existing film in database
func (repo *GormRepository) UpdateFilm(film film.Film, currentVersion int) error {

	filmData := newFilmTable(film)

	err := repo.DB.Model(&filmData).Where("user_id = ?", film.UserID).Updates(FilmTable{Title: filmData.Title, Version: filmData.Version}).Error
	if err != nil {
		return err
	}

	return nil
}
