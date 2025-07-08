package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rickyhuang08/99-backend-exercise/internal/entity"
	"github.com/rickyhuang08/99-backend-exercise/pkg/auth"
)

// CreateListingHandler returns listing
func (h *Handler) CreateListingHandler(c *gin.Context) {
	claims := c.Request.Context().Value(auth.UserKey).(jwt.MapClaims)
	if claims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := claims["user_id"].(float64)
	if !ok || userID <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	var listingPayload entity.ListingPayload
	listingPayload.UserID = int(userID) // Ensure userID is set from JWT claims

	responseListing := ResponseListing{
		Result:  false,
		Listing: entity.Listing{},
		Message: "",
	}

	if err := c.ShouldBind(&listingPayload); err != nil {
		responseListing.Message = "Invalid Request"
		c.JSON(http.StatusBadRequest, responseListing)
		return
	}

	listing, err := h.ListingUsecase.CreateListing(listingPayload)
	if err != nil {
		responseListing.Message = "Failed To Create Listing"
		fmt.Println("error : ", err)
		c.JSON(http.StatusInternalServerError, responseListing)
		return
	}

	responseListing.Result = true
	responseListing.Listing = *listing

	c.JSON(http.StatusOK, responseListing)
}

// GetListing
func (h *Handler) GetListingHandler(c *gin.Context) {
	var params entity.GetListingParam
	responseListings := ResponseGetAllListing{
		Result:  false,
		Listings: []entity.Listing{},
		Message: "",
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		responseListings.Message = "Invalid Request"
		c.JSON(http.StatusBadRequest, responseListings)
		return
	}

	listings, err := h.ListingUsecase.GetListing(params)
	if err != nil {
		responseListings.Message = "Failed to Get Listing"
		fmt.Println("error : ", err)
		c.JSON(http.StatusInternalServerError, responseListings)
		return
	}

	if len(listings) == 1 {
		responseListing := ResponseListing{
			Result:  true,
			Listing: listings[0],
			Message: "",
		}
		c.JSON(http.StatusOK, responseListing)
		return
	}

	responseListings.Result = true
	responseListings.Listings = listings

	c.JSON(http.StatusOK, responseListings)
}

func (h *Handler) GetListingsForPublicHandler(c *gin.Context) {
	var params entity.GetListingParam
	responseListings := ResponseGetAllListingWithUserInfo{
		Result:  false,
		Listings: []entity.ListingWithUserInfo{},
		Message: "",
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		responseListings.Message = "Invalid Request"
		c.JSON(http.StatusBadRequest, responseListings)
		return
	}

	listings, err := h.ListingUsecase.GetListingsForPublic(params)
	if err != nil {
		responseListings.Message = "Failed to Get Listings"
		fmt.Println("error : ", err)
		c.JSON(http.StatusInternalServerError, responseListings)
		return
	}

	if len(listings) == 0 {
		responseListings.Message = "No Listings Found"
		c.JSON(http.StatusNotFound, responseListings)
		return
	}

	responseListings.Result = true
	responseListings.Listings = listings

	c.JSON(http.StatusOK, responseListings)
}