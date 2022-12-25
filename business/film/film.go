package film

import "time"

// Film product film that available to rent or sell
type Film struct {
	ID         int
	UserID     int
	Title      string
	Genre      string
	Artist     string
	Rating     float32
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy string
	Version    int
}

// NewFilm create new film
func NewFilm(
	id int,
	userID int,
	title string,
	genre string,
	artist string,
	rating float32,
	creator string,
	createdAt time.Time) Film {

	return Film{
		ID:         id,
		UserID:     userID,
		Title:      title,
		Genre:      genre,
		Artist:     artist,
		Rating:     rating,
		CreatedAt:  createdAt,
		CreatedBy:  creator,
		ModifiedAt: createdAt,
		ModifiedBy: creator,
		Version:    1,
	}
}

// ModifyFilm update existing film data
func (oldData *Film) ModifyFilm(
	newTitle string,
	// newGenre string,
	// newArtist string,
	// newRating int,
	modifiedAt time.Time,
	updater string) Film {

	return Film{
		ID:     oldData.ID,
		UserID: oldData.UserID,
		Title:  newTitle,
		// Genre: newGenre,
		// Artist: newArtist,
		// Rating: newRating,
		CreatedAt:  oldData.CreatedAt,
		CreatedBy:  oldData.CreatedBy,
		ModifiedAt: modifiedAt,
		ModifiedBy: updater,
		Version:    oldData.Version + 1,
	}
}
