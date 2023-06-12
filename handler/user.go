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
		helper.ErrorResponse(c, "Failed to register the account due to invalid input data.", http.StatusUnprocessableEntity, helper.FormatValidationError(err))
		return
	}

	registeredUser, err := h.userService.RegisterUser(input)
	if err != nil {
		if helper.IsEmailExistsError(err) {
			helper.ErrorResponse(c, "Failed to register the account. The email is already taken.", http.StatusConflict, nil)
			return
		}
		helper.ErrorResponse(c, "Failed to register the account due to internal server error.", http.StatusInternalServerError, nil)
		return
	}

	formatter := user.FormatUser(registeredUser, "TOKENTOKENTOKEN")
	response := helper.APIResponse("Account has been created successfully.", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.ErrorResponse(c, "Failed to login to the account due to invalid input data.", http.StatusUnprocessableEntity, helper.FormatValidationError(err))
		return
	}

	loggedUser, err := h.userService.Login(input)
	if err != nil {
		helper.ErrorResponse(c, "Failed to login to the account. Please check your credentials and try again.", http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	formatter := user.FormatUser(loggedUser, "TOKENTOKENTOKENTOKENTOKENTOKENTOKENTOKENTOKEN")
	response := helper.APIResponse("Login successful.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.ErrorResponse(c, "Failed to check email availability due to invalid input data.", http.StatusUnprocessableEntity, helper.FormatValidationError(err))
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		helper.ErrorResponse(c, "Failed to check email availability.", http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email is not available"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
