package handler

import (
	"net/http"
	"startup/helper"
	"startup/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Account Failed Created", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Account Failed Created", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// token,err := h.jwtService.GenerateToken

	formatter := user.FormatUser(newUser, "halo")

	response := helper.ApiResponse("Account Has Ben Created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
