package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	user "project_api/User"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

//		@Summary		Login User
//		@Description	User Login
//	 	@Schemes
//		@Tags			auth
//		@Accept			json
//		@Param			create	body	user.InputUser	true	"User login"
//		@Produce		json
//		@Success		200		{object} 	Response
//		@Security		ApiKeyAuth
//		@Router			/auth/login [post]
func (h *userHandler) Login(c *gin.Context) {
	var input user.InputUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Verify your data format or structure",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.userService.LoginCheck(input)
	if err != nil {
		response := gin.H{
			"Success": false,
			"Message": "Password / username is incorrect",
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := gin.H{
		"Success": true,
		"Message": "Login confirmed",
		"token":   token,
	}

	c.JSON(http.StatusOK, response)
}

//		@Summary		Register User
//		@Description	User Register
//	 	@Schemes
//		@Tags			auth
//		@Accept			json
//		@Param			create	body	user.InputUser	true	"User register"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/auth/register [post]
func (h *userHandler) Register(c *gin.Context) {
	var input user.InputUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Verify your data format or structure",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = h.userService.Register(input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := gin.H{
		"Success": true,
		"Message": "new Account successfully added",
	}

	c.JSON(http.StatusOK, response)
}
