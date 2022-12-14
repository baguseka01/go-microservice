package user

import (
	"time"

	"github.com/baguseka01/golang_microservice_hexagonal/business"
	"github.com/baguseka01/golang_microservice_hexagonal/util/validator"
)

// InsertUserSpec
type InsertUserSpec struct {
	Name     string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}

// <====================================== THE IMPLEMENTATION OF THOSE INTERFACE PUT BELOW ======================================>

type service struct {
	repository Repository
}

// NewService Construct user service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

// FindUserByID Get user by given ID, return nil if not exis
func (s *service) FindUserByID(id int) (*User, error) {
	return s.repository.FindUserByID(id)
}

// FindUserByUsernameAndPassword Get user by given ID, return nil if not exis
func (s *service) FindUserByUsernameAndPassword(username string, password string) (*User, error) {
	return s.repository.FindUserByUsernameAndPassword(username, password)
}

// FindAllUser Get all users, will be return empty array if no data or error occured
func (s *service) FindAllUser(skip int, rowPerPage int) ([]User, error) {

	user, err := s.repository.FindAllUser(skip, rowPerPage)
	if err != nil {
		return []User{}, err
	}

	return user, err
}

// InsertUser Create new user and store into database
func (s *service) InsertUser(insertUserSpec InsertUserSpec, createdBy string) error {
	err := validator.GetValidator().Struct(insertUserSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	user := NewUser(
		1,
		insertUserSpec.Name,
		insertUserSpec.Username,
		insertUserSpec.Password,
		createdBy,
		time.Now(),
	)

	err = s.repository.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser will update found user, if not exists will be return error
func (s *service) UpdateUser(id int, name string, midifiedBy string, currentVersion int) error {

	user, err := s.repository.FindUserByID(id)
	if err != nil {
		return err
	} else if user == nil {
		return business.ErrNotFound
	} else if user.Version != currentVersion {
		return business.ErrHasBeenModified
	}

	modifiedUser := user.ModifyUser(name, time.Now(), midifiedBy)

	return s.repository.UpdateUser(modifiedUser, currentVersion)
}
