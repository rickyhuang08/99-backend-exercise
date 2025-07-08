package usecase

import (
	"errors"

	"github.com/rickyhuang08/99-backend-exercise/helpers"
	"github.com/rickyhuang08/99-backend-exercise/internal/entity"
	"github.com/rickyhuang08/99-backend-exercise/internal/repository/sql"
	"golang.org/x/crypto/bcrypt"
)

// UserUsecase handles user logic
type UserUsecase struct {
	UserRepo *sql.UserRepository
}

// NewUserUsecase initializes user usecase
func NewUserUsecase(userRepo *sql.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (uc *UserUsecase) CreateUser(userPayload entity.UserPayload) (*entity.User, error) {
	// Here we would typically hash the password before saving
	// For simplicity, we are not doing that in this mock implementation
	if userPayload.Name == "" || userPayload.Email == "" || userPayload.Password == "" {
		return nil, errors.New(helpers.ErrInvalidUserPayload)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPayload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// Set the hashed password in the user payload
	userPayload.Password = string(hashedPassword)
	return uc.UserRepo.Create(userPayload)
}

func (uc *UserUsecase) GetUser(params entity.GetUserParam) ([]entity.User, error) {
	// Fetch user from repository
	return uc.UserRepo.GetUser(params)
}