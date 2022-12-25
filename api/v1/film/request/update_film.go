package request

// UpdateFilmRequest update Film request paylod
type UpdateFilmRequest struct {
	Title   string `json:"title"`
	Version int    `json:"version"`
}
