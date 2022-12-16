package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseModel struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Ok      bool        `json:"ok"`
}

func HandleSuccess(c *gin.Context, data interface{}) {
	responData := ResponseModel{
		Status:  "200",
		Message: "Success",
		Data:    data,
		Ok:      true,
	}
	c.JSON(http.StatusOK, responData)
}

func HandleCreated(c *gin.Context, data interface{}) {
	responData := ResponseModel{
		Status:  "201",
		Message: "Success",
		Data:    data,
		Ok:      true,
	}
	c.JSON(http.StatusCreated, responData)
}

func HandleError(c *gin.Context, status int, message string) {
	responData := ResponseModel{
		Status:  strconv.Itoa(status),
		Message: message,
		Ok:      false,
	}
	c.JSON(status, responData)
}
