package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

func SendDefaultBadRequestError(ctx *gin.Context) {
	SendBadRequestError(ctx, gin.H{
		"errorMessage": "Bad request",
	})
}

func SendBadRequestError(ctx *gin.Context, obj any) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, obj)
}

func SendValidationError(ctx *gin.Context, err error) {
	var message []string
	if validationErrors, ok := err.(v.ValidationErrors); ok {
		for _, value := range validationErrors {
			message = append(message, value.Error())
		}
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
	return
}
