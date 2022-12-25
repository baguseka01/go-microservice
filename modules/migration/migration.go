package migration

import (
	"github.com/baguseka01/golang_microservice_hexagonal/modules/film"
	"github.com/baguseka01/golang_microservice_hexagonal/modules/user"
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.UserTable{}, &film.FilmTable{})
}
