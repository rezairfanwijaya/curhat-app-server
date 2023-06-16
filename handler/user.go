package handler

import (
	"fly/helper"
	"fly/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handerUser struct {
	serviceUser user.IService
}

func NewHandlerUser(serviceUser user.IService) *handerUser {
	return &handerUser{serviceUser}
}

func (h *handerUser) Create(c *gin.Context) {
	var input user.InputNewUser

	if err := c.BindJSON(&input); err != nil {
		errBinding := helper.GenerateErrorBinding(err)
		response := helper.GenerateResponseAPI(
			http.StatusBadRequest,
			"error",
			errBinding,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userSaved, httpCode, err := h.serviceUser.Save(input)
	if err != nil {
		response := helper.GenerateResponseAPI(
			httpCode,
			"error",
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	userFormatted := user.FormatterUser(userSaved)
	response := helper.GenerateResponseAPI(
		httpCode,
		"success",
		userFormatted,
	)

	c.JSON(httpCode, response)
}

func (h *handerUser) GetByEmail(c *gin.Context) {
	email := c.Param("email")

	userByEmail, httpCode, err := h.serviceUser.GetByEmail(email)
	if err != nil {
		response := helper.GenerateResponseAPI(
			httpCode,
			"error",
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	userFormatted := user.FormatterUser(userByEmail)
	response := helper.GenerateResponseAPI(
		httpCode,
		"success",
		userFormatted,
	)

	c.JSON(httpCode, response)
}
