package response

import (
	"time"

	"github.com/baguseka01/golang_microservice_hexagonal/business/user"
)

// GetUserResponse Get user by ID response payload
type GetUserResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	ModifiedAt time.Time `json:"modified_at"`
	Version    int       `json:"version"`
}

// NewGetUserResponse construct GetUserResponsef
func NewGetUserResponse(user user.User) *GetUserResponse {
	var getUserResponse GetUserResponse

	getUserResponse.ID = user.ID
	getUserResponse.Name = user.Name
	getUserResponse.Username = user.Username
	getUserResponse.ModifiedAt = user.ModifiedAt
	getUserResponse.Version = user.Version

	return &getUserResponse
}
