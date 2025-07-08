package entity

type Listing struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Price       int    `json:"price"`
	ListingType string `json:"listing_type"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type ListingWithUserInfo struct {
	ID          int    `json:"id"`
	Price       int    `json:"price"`
	ListingType string `json:"listing_type"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	User        User   `json:"user"`
}

type ListingPayload struct {
	UserID      int    `form:"user_id"`
	Price       int    `form:"price"`
	ListingType string `form:"listing_type"`
}

type GetListingParam struct {
	ID       int `form:"id"`
	UserID   int `form:"user_id"`
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}
