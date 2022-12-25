package user_test

import (
	"os"
	"testing"
	"time"

	"github.com/baguseka01/golang_microservice_hexagonal/business"
	"github.com/baguseka01/golang_microservice_hexagonal/business/user"
	userMock "github.com/baguseka01/golang_microservice_hexagonal/business/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id       = 1
	name     = "name"
	username = "username"
	password = "password"
	creator  = "creator"

	modifier = "modifier"
	version  = 1
)

var (
	userService    user.Service
	userRepository userMock.Repository

	userData       user.User
	insertUserData user.InsertUserSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindUserByID(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()

		user, err := userService.FindUserByID(id)

		assert.Nil(t, err)

		assert.NotNil(t, user)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)
	})

	t.Run("Expect user not found", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		user, err := userService.FindUserByID(id)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func setup() {
	userData = user.NewUser(
		id,
		name,
		username,
		password,
		creator,
		time.Now(),
	)

	insertUserData = user.InsertUserSpec{
		Name:     name,
		Username: username,
		Password: password,
	}

	userService = user.NewService(&userRepository)
}
