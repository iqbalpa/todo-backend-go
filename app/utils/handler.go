package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, message string, data interface{}) {
	response := Response{
		Status: "SUCCESS",
		Message: message,
		Data: data,
	}
	ctx.JSON(http.StatusOK, response)
}

func FailedResponse(ctx *gin.Context, status string, message string, data interface{}) {
	response := Response{
		Status: status,
		Message: message,
		Data: data,
	}
	switch status {
	case "FORBIDDEN":
		ctx.JSON(http.StatusForbidden, response)
	case "INTERNAL SERVER ERROR":
		ctx.JSON(http.StatusInternalServerError, response)
	default:
		ctx.JSON(http.StatusBadRequest, response)
	}
}