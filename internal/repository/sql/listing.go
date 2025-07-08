package sql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rickyhuang08/99-backend-exercise/helpers"
	"github.com/rickyhuang08/99-backend-exercise/internal/entity"
)

// ListingRepository handles DB interactions
type ListingRepository struct {
	DB *sql.DB
	TimeModule helpers.TimeProvider
}

// NewListingRepository initializes the repo with dummy data
func NewListingRepository(db *sql.DB, timeModule helpers.TimeProvider) *ListingRepository {
	return &ListingRepository{DB: db, TimeModule: timeModule}
}

func (r *ListingRepository) Create(listingPayload entity.ListingPayload) (*entity.Listing, error) {
	now := r.TimeModule.Now().UnixMicro()

	// Insert listing into the database
	res, err := r.DB.Exec("INSERT INTO listings (user_id, price, listing_type, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		listingPayload.UserID, listingPayload.Price, listingPayload.ListingType, now, now)

	if err != nil {
		return nil, fmt.Errorf("failed to create listing: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %w", err)
	}
	listing := entity.Listing{
		ID:          int(id),
		UserID:      listingPayload.UserID,
		Price:       listingPayload.Price,
		ListingType: listingPayload.ListingType,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return &listing, nil
}

// GetListings retrieves listings by ID or user_id optional parameter
// If both are empty, it returns an error.
// If ID is provided, it fetches by ID; if user_id is provided, it fetches by user_id.
func (r *ListingRepository) GetListings(params entity.GetListingParam) ([]entity.Listing, error) {
	// Validate pagination input
	if params.Page <= 0 {
		params.Page = helpers.DefaultPage
	}
	if params.PageSize <= 0 {
		params.PageSize = helpers.DefaultPageSize
	}
	offset := (params.Page - 1) * params.PageSize

	var listings []entity.Listing
	query := "SELECT id, user_id, price, listing_type, created_at, updated_at FROM listings"
	args := []interface{}{}

	if params.ID > 0 || params.UserID > 0 {
		query += " WHERE"
	}
	
	if params.ID > 0 {
		query += " id = ?"
		args = append(args, params.ID)
	} else if params.UserID > 0 {
		query += " user_id = ?"
		args = append(args, params.UserID)
	}

	if params.ID > 0 {
		query += " LIMIT 1"
	} else {
		query += " ORDER BY created_at DESC LIMIT ? OFFSET ?"
		args = append(args, params.PageSize, offset)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var listing entity.Listing
		if err := rows.Scan(&listing.ID, &listing.UserID, &listing.Price, &listing.ListingType, &listing.CreatedAt, &listing.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan listing: %w", err)
		}
		listings = append(listings, listing)
	}

	if len(listings) == 0 {
		return nil, errors.New("no listings found")
	}

	return listings, nil
}


// GetListings retrieves listings by ID or user_id optional parameter
// If both are empty, it returns an error.
// If ID is provided, it fetches by ID; if user_id is provided, it fetches by user_id.
func (r *ListingRepository) GetListingsForPublic(params entity.GetListingParam) ([]entity.ListingWithUserInfo, error) {
	// Validate pagination input
	if params.Page <= 0 {
		params.Page = helpers.DefaultPage
	}
	if params.PageSize <= 0 {
		params.PageSize = helpers.DefaultPageSize
	}
	offset := (params.Page - 1) * params.PageSize

	var listings []entity.ListingWithUserInfo
	query := `
		SELECT 
			l.id, l.price, l.listing_type, l.created_at, l.updated_at,
			u.id, u.name, u.email, u.created_at, u.updated_at
		FROM listings l
		JOIN users u ON l.user_id = u.id
	`
	args := []interface{}{}

	if params.ID > 0 || params.UserID > 0 {
		query += " WHERE"
	}
	
	if params.ID > 0 {
		query += " l.id = ?"
		args = append(args, params.ID)
	} else if params.UserID > 0 {
		query += " l.user_id = ?"
		args = append(args, params.UserID)
	}

	if params.ID > 0 {
		query += " LIMIT 1"
	} else {
		query += " ORDER BY l.created_at DESC LIMIT ? OFFSET ?"
		args = append(args, params.PageSize, offset)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var listing entity.ListingWithUserInfo
		if err := rows.Scan(
			&listing.ID, 
			&listing.Price, 
			&listing.ListingType, 
			&listing.CreatedAt, 
			&listing.UpdatedAt,
			&listing.User.ID,
			&listing.User.Name,
			&listing.User.Email,
			&listing.User.CreatedAt,
			&listing.User.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan listing: %w", err)
		}
		listings = append(listings, listing)
	}

	if len(listings) == 0 {
		return nil, errors.New("no listings found")
	}

	return listings, nil
}