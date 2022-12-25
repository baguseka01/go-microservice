package auth

import (
	"os"
	"time"

	"github.com/baguseka01/golang_microservice_hexagonal/business/user"
	"github.com/golang-jwt/jwt"
)

// <====================================== THE IMPLEMENTATION OF THOSE INTERFACE PUT BELOW ======================================>

type service struct {
	userService user.Service
}

func NewService(userService user.Service) Service {
	return &service{
		userService,
	}
}

// Login by given user Username and Password, return error if not exist
func (s *service) Login(username string, password string) (string, error) {
	
	user, err := s.userService.FindUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires after 1 hour
	claims["name"] = user.Name

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}