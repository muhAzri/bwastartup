package user

import (
	"errors"
	"strings"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	user := User{
		ID:         uuid.NewString(),
		Name:       input.Name,
		Email:      input.Email,
		Occupation: input.Occupation,
		Role:       "user",
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func IsEmailExistsError(err error) bool {
	var ErrEmailExists = errors.New("email already exists")

	if err == ErrEmailExists {
		return true
	}

	if err != nil {
		// Check the error message or code to identify the "email already exists" error
		errMsg := err.Error()
		if strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			return true
		}
		if strings.Contains(errMsg, "23505") { // Assuming SQLSTATE 23505 represents the "unique constraint violation" error
			return true
		}
	}

	return false
}
