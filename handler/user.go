package handler

import (
	"fly/helper"
	"fly/user"
	"net/http"
	"strconv"

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

func (h *handerUser) GetAll(c *gin.Context) {
	users, httpCode, err := h.serviceUser.GetAll()
	if err != nil {
		response := helper.GenerateResponseAPI(
			httpCode,
			"error",
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	usersFormatted := user.FormatterUsers(users)
	response := helper.GenerateResponseAPI(
		httpCode,
		"success",
		usersFormatted,
	)

	c.JSON(httpCode, response)
}

func (h *handerUser) Update(c *gin.Context) {
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

	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil || idNumber <= 0 {
		response := helper.GenerateResponseAPI(
			http.StatusBadRequest,
			"error",
			"id harus berupa integer dan lebih besar dari 0",
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userUpdated, httpCode, err := h.serviceUser.Update(input, idNumber)
	if err != nil {
		response := helper.GenerateResponseAPI(
			httpCode,
			"error",
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	userFormatted := user.FormatterUser(userUpdated)
	response := helper.GenerateResponseAPI(
		httpCode,
		"success",
		userFormatted,
	)

	c.JSON(httpCode, response)
}

func (h *handerUser) Delete(c *gin.Context) {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil || idNumber <= 0 {
		response := helper.GenerateResponseAPI(
			http.StatusBadRequest,
			"error",
			"id harus berupa integer dan lebih besar dari 0",
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	httpCode, err := h.serviceUser.Delete(idNumber)
	if err != nil {
		response := helper.GenerateResponseAPI(
			httpCode,
			"error",
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	response := helper.GenerateResponseAPI(
		httpCode,
		"success",
		"success",
	)

	c.JSON(httpCode, response)
}
