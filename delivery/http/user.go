package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rickyhuang08/99-backend-exercise/internal/entity"
)

// CreateUserHandler returns user
func (h *Handler) CreateUserHandler(c *gin.Context) {
	var userPayload entity.UserPayload
	responseUser := ResponseUser{
		Result:  false,
		User:    entity.User{},
		Message: "",
	}

	if err := c.ShouldBind(&userPayload); err != nil {
		responseUser.Message = "Invalid Request"
		c.JSON(http.StatusBadRequest, responseUser)
		return
	}

	user, err := h.UserUsecase.CreateUser(userPayload)
	if err != nil {
		responseUser.Message = "Failed To Create User"
		fmt.Println("error : ", err)
		c.JSON(http.StatusInternalServerError, responseUser)
		return
	}

	responseUser.Result = true
	responseUser.User = *user

	c.JSON(http.StatusOK, responseUser)
}

// GetUser
func (h *Handler) GetUserHandler(c *gin.Context) {
	var params entity.GetUserParam
	responseUsers := ResponseGetAllUser{
		Result:  false,
		Users:   []entity.User{},
		Message: "",
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		responseUsers.Message = "Invalid Request"
		c.JSON(http.StatusBadRequest, responseUsers)
		return
	}

	users, err := h.UserUsecase.GetUser(params)
	if err != nil {
		responseUsers.Message = "Failed to Get User"
		fmt.Println("error : ", err)
		c.JSON(http.StatusInternalServerError, responseUsers)
		return
	}

	if len(users) == 1 {
		responseUser := ResponseUser{
			Result:  true,
			User:    users[0],
			Message: "",
		}
		c.JSON(http.StatusOK, responseUser)
		return
	}

	responseUsers.Result = true
	responseUsers.Users = users

	c.JSON(http.StatusOK, responseUsers)
}
