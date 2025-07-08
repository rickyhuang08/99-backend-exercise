package usecase

import (
	"errors"

	"github.com/rickyhuang08/99-backend-exercise/helpers"
	"github.com/rickyhuang08/99-backend-exercise/internal/entity"
	"github.com/rickyhuang08/99-backend-exercise/internal/repository/sql"
	"github.com/rickyhuang08/99-backend-exercise/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

// AuthUsecase handles authentication
type AuthUsecase struct {
	UserRepo  *sql.UserRepository
	JwtHelper *auth.JWTHelper
}

// NewAuthUsecase initializes auth usecase
func NewAuthUsecase(userRepo *sql.UserRepository, jwtHelper *auth.JWTHelper) *AuthUsecase {
	return &AuthUsecase{
		UserRepo:  userRepo,
		JwtHelper: jwtHelper,
	}
}

// Login checks user credentials (dummy check for now)
func (uc *AuthUsecase) Login(req entity.LoginRequest) (string, error) {
	user, err := uc.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New(helpers.UserNotFoundError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New(helpers.InvalidCredentialsError)
	}

	token, err := uc.JwtHelper.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New(helpers.FailedToGenerateTokenError)
	}

	return token, nil
}
