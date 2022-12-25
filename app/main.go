package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/baguseka01/golang_microservice_hexagonal/api"
	userController "github.com/baguseka01/golang_microservice_hexagonal/api/v1/user"
	userService "github.com/baguseka01/golang_microservice_hexagonal/business/user"
	migration "github.com/baguseka01/golang_microservice_hexagonal/modules/migration"
	userRepository "github.com/baguseka01/golang_microservice_hexagonal/modules/user"

	filmController "github.com/baguseka01/golang_microservice_hexagonal/api/v1/film"
	filmService "github.com/baguseka01/golang_microservice_hexagonal/business/film"
	filmRepository "github.com/baguseka01/golang_microservice_hexagonal/modules/film"

	authController "github.com/baguseka01/golang_microservice_hexagonal/api/v1/auth"
	authService "github.com/baguseka01/golang_microservice_hexagonal/business/auth"

	"github.com/baguseka01/golang_microservice_hexagonal/config"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {

	configDB := map[string]string{
		"DB_Username": os.Getenv("GO_DB_USERNAME"),
		"DB_Password": os.Getenv("GO_DB_PASSWORD"),
		"DB_Port":     os.Getenv("GO_DB_PORT"),
		"DB_Host":     os.Getenv("GO_DB_ADDRESS"),
		"DB_Name":     os.Getenv("GO_DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])
	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db
}

func main() {
	// load config if available or set to default
	config := config.GetConfig()

	// initialize database connection based on given config
	dbConnection := newDatabaseConnection(config)

	// initiate user repository
	userRepo := userRepository.NewGormDBRepository(dbConnection)

	// initeate user service
	userService := userService.NewService(userRepo)

	// initiate user controller
	userController := userController.NewController(userService)

	// initiate film repository
	filmRepo := filmRepository.NewGormDBRepository(dbConnection)

	// initiate film service
	filmService := filmService.NewService(filmRepo)

	// initiate film controller
	filmController := filmController.NewController(filmService)

	// initiate auth service
	authService := authService.NewService(userService)

	// initiate auth controller
	authController := authController.NewController(authService)

	// create acho http
	e := echo.New()

	// register API path and hendler
	api.RegisterPath(e, authController, userController, filmController)

	// run server
	go func() {
		address := fmt.Sprintf("localhost:%d", config.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 second to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
