package user

import (
	"strconv"

	"github.com/baguseka01/golang_microservice_hexagonal/api/common"
	"github.com/baguseka01/golang_microservice_hexagonal/api/paginator"
	"github.com/baguseka01/golang_microservice_hexagonal/api/v1/user/request"
	"github.com/baguseka01/golang_microservice_hexagonal/api/v1/user/response"
	"github.com/baguseka01/golang_microservice_hexagonal/business/user"
	echo "github.com/labstack/echo/v4"
)

// Controller Get user API controller
type Controller struct {
	service user.Service
}

// NewController Construct user API controller
func NewController(service user.Service) *Controller {
	return &Controller{
		service,
	}
}

// FindUserByID Find user by ID echo handler
func (controller *Controller) FindUserByID(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))

	user, err := controller.service.FindUserByID(id)
	if err != nil {
		return e.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetUserResponse(*user)

	return e.JSON(common.NewSuccessResponse(response))
}

// FindAllUser Find All User with pagination handler
func (controller *Controller) FindAllUser(e echo.Context) error {
	pageQueryParam := e.QueryParam("page")
	rowPerPageQueryParam := e.QueryParam("row_per_page")

	skip, page, rowPerPage := paginator.CreatePagination(pageQueryParam, rowPerPageQueryParam)

	users, err := controller.service.FindAllUser(skip, rowPerPage)
	if err != nil {
		return e.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAllUserRespponse(users, page, rowPerPage)

	return e.JSON(common.NewSuccessResponse(response))
}

// InsertUser Create new user handler
func (controller *Controller) InsertUser(e echo.Context) error {
	insertUserRequest := new(request.InsertUserRequest)

	if err := e.Bind(insertUserRequest); err != nil {
		return e.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.InsertUser(*insertUserRequest.ToUpsertUserSpec(), "creator")
	if err != nil {
		return e.JSON(common.NewErrorBusinessResponse(err))
	}

	return e.JSON(common.NewSuccessResponseWithoutData())
}

// UpdateUser update existing user handler
func (controller *Controller) UpdateUser(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))

	updateUserRequest := new(request.UpdateUserRequest)

	if err := e.Bind(updateUserRequest); err != nil {
		return e.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdateUser(id, updateUserRequest.Name, "modifier", updateUserRequest.Version)
	if err != nil {
		e.JSON(common.NewErrorBusinessResponse(err))
	}

	return e.JSON(common.NewSuccessResponseWithoutData())
}
