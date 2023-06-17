package handler

import (
	"fly/helper"
	"fly/note"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerNote struct {
	serviceNote note.IService
}

func NewHandlerNote(serviceNote note.IService) *handlerNote {
	return &handlerNote{serviceNote}
}

func (h *handlerNote) Create(c *gin.Context) {
	var input note.InputNewNote

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

	noteSaved, httpCode, err := h.serviceNote.Save(input)
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
		noteSaved,
	)

	c.JSON(httpCode, response)
}

func (h *handlerNote) GetAll(c *gin.Context) {
	notes, httpCode, err := h.serviceNote.GetAll()
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
		notes,
	)

	c.JSON(httpCode, response)
}

func (h *handlerNote) Delete(c *gin.Context) {
	httpCode, err := h.serviceNote.Delete()
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
