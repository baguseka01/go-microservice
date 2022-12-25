package film

import (
	"strconv"

	"github.com/baguseka01/golang_microservice_hexagonal/api/common"
	"github.com/baguseka01/golang_microservice_hexagonal/api/paginator"
	"github.com/baguseka01/golang_microservice_hexagonal/api/v1/film/request"
	"github.com/baguseka01/golang_microservice_hexagonal/api/v1/film/response"
	"github.com/baguseka01/golang_microservice_hexagonal/business/film"
	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
)

// Controller Get film API controller
type Controller struct {
	service film.Service
}

// NewController Construct film API controller
func NewController(service film.Service) *Controller {
	return &Controller{
		service,
	}
}

// FindFilmByID Find film by ID echo handler
func (controller *Controller) FindFilmByID(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	if !user.Valid {
		return e.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	// use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	// MUST CONVERT TO FLOAT^$< OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return e.JSON(common.NewForbiddenResponse())
	}

	id, _ := strconv.Atoi(e.Param("id"))

	film, err := controller.service.FindFilmByID(id, int(userID))
	if err != nil {
		return e.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetFilmResponse(*film)

	return e.JSON(common.NewSuccessResponse(response))
}

// FindAllFilm Find All Film with pagination handler
func (controller *Controller) FindAllFilm(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	if !user.Valid {
		return e.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	// use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	// MUST CONVERT TO FLOAT^$< OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return e.JSON(common.NewForbiddenResponse())
	}

	pageQueryParam := e.QueryParam("page")
	rowPerPageQueryParam := e.QueryParam("row_per_page")

	skip, page, rowPerPage := paginator.CreatePagination(pageQueryParam, rowPerPageQueryParam)

	films, err := controller.service.FindAllFilm(int(userID), skip, rowPerPage)
	if err != nil {
		return e.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAllFilmRespponse(films, page, rowPerPage)

	return e.JSON(common.NewSuccessResponse(response))

}

// InsertFilm Create new film echo handler
func (controller *Controller) InsertFilm(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	if !user.Valid {
		return e.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	// use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	// MUST CONVERT TO FLOAT^$< OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return e.JSON(common.NewForbiddenResponse())
	}

	insertFilmRequest := new(request.InsertFilmRequest)
	if err := e.Bind(insertFilmRequest); err != nil {
		return e.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.InsertFilm(*insertFilmRequest.ToUpsertFilmSpec(int(userID)), "creator")
	if err != nil {
		return e.JSON(common.NewErrorBusinessResponse(err))
	}

	return e.JSON(common.NewSuccessResponseWithoutData())
}

// UpdateFilm update existing film
func (controller *Controller) UpdateFilm(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	if !user.Valid {
		return e.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	// use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	// MUST CONVERT TO FLOAT^$< OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return e.JSON(common.NewForbiddenResponse())
	}

	id, _ := strconv.Atoi(e.Param("id"))

	updateFilmRequest := new(request.UpdateFilmRequest)
	if err := e.Bind(updateFilmRequest); err != nil {
		return e.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdateFilm(id, int(userID), updateFilmRequest.Title, "modifier", updateFilmRequest.Version)
	if err != nil {
		return e.JSON(common.NewErrorBusinessResponse(err))
	}

	return e.JSON(common.NewSuccessResponseWithoutData())
}
