package http

import "github.com/rickyhuang08/99-backend-exercise/internal/entity"

type ResponseUser struct {
	Result  bool        `json:"result"`
	User    entity.User `json:"user"`
	Message string      `json:"message"`
}

type ResponseGetAllUser struct {
	Result  bool          `json:"result"`
	Users   []entity.User `json:"users"`
	Message string        `json:"message"`
}

type ResponseListing struct {
	Result  bool           `json:"result"`
	Listing entity.Listing `json:"listing"`
	Message string         `json:"message"`
}

type ResponseGetAllListing struct {
	Result   bool             `json:"result"`
	Listings []entity.Listing `json:"listings"`
	Message  string           `json:"message"`
}

type ResponseGetAllListingWithUserInfo struct {
	Result   bool                         `json:"result"`
	Listings []entity.ListingWithUserInfo `json:"listings"`
	Message  string                       `json:"message"`
}
