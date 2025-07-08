package http

import "github.com/rickyhuang08/99-backend-exercise/internal/usecase"

// Handler struct contains all use cases
type Handler struct {
	AuthUsecase *usecase.AuthUsecase
	UserUsecase *usecase.UserUsecase
	ListingUsecase *usecase.ListingUsecase
}

// NewHandler initializes an HTTP handler
func NewHandler(authUC *usecase.AuthUsecase, userUC *usecase.UserUsecase, listingUC *usecase.ListingUsecase) *Handler {
	return &Handler{AuthUsecase: authUC, UserUsecase: userUC, ListingUsecase: listingUC}
}