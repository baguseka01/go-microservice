package response

import (
	"github.com/baguseka01/golang_microservice_hexagonal/api/paginator"
	"github.com/baguseka01/golang_microservice_hexagonal/business/user"
)

type getAllUserResponse struct {
	Meta  paginator.Meta `json:"meta"`
	Users []GetUserResponse
}

// NewGetAllUserRespponse construct GetAllUserResponse
func NewGetAllUserRespponse(users []user.User, page int, roePerPage int) getAllUserResponse {
	var (
		lenUsers = len(users)
	)

	getAllUserResponse := getAllUserResponse{}
	getAllUserResponse.Meta.BuildMeta(lenUsers, page, roePerPage)

	for index, value := range users {
		if index == getAllUserResponse.Meta.Page {
			continue
		}

		var getUserResponse GetUserResponse

		getUserResponse.ID = value.ID
		getUserResponse.Name = value.Name
		getUserResponse.Username = value.Username
		getUserResponse.ModifiedAt = value.ModifiedAt
		getUserResponse.Version = value.Version

		getAllUserResponse.Users = append(getAllUserResponse.Users, getUserResponse)
	}

	if getAllUserResponse.Users == nil {
		getAllUserResponse.Users = []GetUserResponse{}
	}

	return getAllUserResponse
}
