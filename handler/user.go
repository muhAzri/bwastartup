package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Failed to register account. Invalid input data.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	registeredUser, err := h.userService.RegisterUser(input)
	if err != nil {
		if helper.IsEmailExistsError(err) {
			c.JSON(http.StatusConflict, helper.APIResponse("ailed to register account. Email is already taken.", http.StatusConflict, "error", nil))
			return
		}
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to register account. Internal Server error.", http.StatusInternalServerError, "error", nil))
		return
	}

	formatter := user.FormatUser(registeredUser, "TOKENTOKENTOKEN")
	response := helper.APIResponse("Account has been created successfully", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		response := helper.APIResponse("Failed to login account. Invalid input data.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to login account", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedUser, "TOKENTOKENTOKENTOKENTOKENTOKENTOKENTOKENTOKEN")
	response := helper.APIResponse("Login successfully", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusOK, response)
}
