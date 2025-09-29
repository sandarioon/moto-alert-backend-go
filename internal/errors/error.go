package errors

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)

	var errMsg string
	switch statusCode {
	case 400:
		errMsg = "BadRequestException"
	case 401:
		errMsg = "UnauthorizedException"
	case 403:
		errMsg = "ForbiddenException"
	default:
		errMsg = "InternalServerError"
	}
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Status: statusCode, Error: errMsg, Message: message})
}
