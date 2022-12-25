package api

import (
	"github.com/baguseka01/golang_microservice_hexagonal/api/middleware"
	"github.com/baguseka01/golang_microservice_hexagonal/api/v1/auth"
	"github.com/baguseka01/golang_microservice_hexagonal/api/v1/film"
	"github.com/baguseka01/golang_microservice_hexagonal/api/v1/user"
	echo "github.com/labstack/echo/v4"
)

// RegisterPath Register all API with routing path
func RegisterPath(e *echo.Echo, authController *auth.Controller, userController *user.Controller, filmController *film.Controller) {
	if authController == nil || userController == nil || filmController == nil {
		panic("Controller parameter cannot be nil")
	}

	// authentication with Versioning endpoint
	authV1 := e.Group("v1/auth")
	authV1.POST("/login", authController.Login)

	// user with Versioning endpoint
	userV1 := e.Group("v1/users")
	userV1.GET("/:id", userController.FindUserByID)
	userV1.GET("", userController.FindAllUser)
	userV1.POST("", userController.InsertUser)
	userV1.PUT("/:id", userController.UpdateUser)

	// film with Versioning endpoint
	filmV1 := e.Group("v1/films")
	filmV1.Use(middleware.JWTMiddleware())
	filmV1.GET("/:id", filmController.FindFilmByID)
	filmV1.GET("", filmController.FindAllFilm)
	filmV1.POST("", filmController.InsertFilm)
	filmV1.PUT("/:id", filmController.UpdateFilm)

	// health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

}
