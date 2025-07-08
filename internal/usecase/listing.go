package usecase

import (
	"errors"

	"github.com/rickyhuang08/99-backend-exercise/helpers"
	"github.com/rickyhuang08/99-backend-exercise/internal/entity"
	"github.com/rickyhuang08/99-backend-exercise/internal/repository/sql"
)

// ListingUsecase handles listing logic
type ListingUsecase struct {
	ListingRepo *sql.ListingRepository
}

// NewListingUsecase initializes listing usecase
func NewListingUsecase(listingRepo *sql.ListingRepository) *ListingUsecase {
	return &ListingUsecase{ListingRepo: listingRepo}
}

func (uc *ListingUsecase) CreateListing(listingPayload entity.ListingPayload) (*entity.Listing, error) {
	// Here we would typically validate the listing payload
	if listingPayload.UserID == 0 || listingPayload.Price <= 0 {
		return nil, errors.New(helpers.ErrInvalidListingPayload)
	}
	return uc.ListingRepo.Create(listingPayload)
}

func (uc *ListingUsecase) GetListing(params entity.GetListingParam) ([]entity.Listing, error) {
	// Fetch listing from repository
	return uc.ListingRepo.GetListings(params)
}

func (uc *ListingUsecase) GetListingsForPublic(params entity.GetListingParam) ([]entity.ListingWithUserInfo, error) {
	// Fetch listing from repository
	listings, err := uc.ListingRepo.GetListingsForPublic(params)
	if err != nil {
		return nil, err
	}
	if len(listings) == 0 {
		return nil, errors.New("listing not found")
	}
	return listings, nil
}