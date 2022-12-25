package auth

import (
	"github.com/baguseka01/golang_microservice_hexagonal/api/common"
	"github.com/baguseka01/golang_microservice_hexagonal/api/v1/auth/request"
	"github.com/baguseka01/golang_microservice_hexagonal/api/v1/auth/response"
	"github.com/baguseka01/golang_microservice_hexagonal/business/auth"

	echo "github.com/labstack/echo/v4"
)

// Controller Get item API controller
type Controller struct {
	service auth.Service
}

// NewController Construct item Api controller
func NewController(service auth.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) Login(e echo.Context) error {
	loginRequest := new(request.LoginRequest)

	if err := e.Bind(loginRequest); err != nil {
		return e.JSON(common.NewBadRequestResponse())
	}

	token, err := controller.service.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return e.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewLoginResponse(token)

	return e.JSON(common.NewSuccessResponse(response))
}
