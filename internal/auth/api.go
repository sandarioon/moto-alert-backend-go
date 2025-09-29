package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sandarioon/moto-alert-backend-go/internal/errors"
	"github.com/sandarioon/moto-alert-backend-go/models"
	"github.com/sandarioon/moto-alert-backend-go/models/dto"
)

type resource struct {
	service Service
}

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *gin.RouterGroup, service Service) {
	res := resource{service}

	r.POST("/create", res.createUser)
	r.POST("/verifyCode", res.verifyCode)
	r.POST("/verifyEmail", res.verifyEmail)
	r.POST("/forgotPassword", res.forgotPassword)
	r.POST("/login", res.login)

}

func (r resource) createUser(c *gin.Context) {
	var input dto.CreateUserRequest

	time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	_, err := r.service.CreateUser(ctx, input)
	if err != nil {
		if (err.Error() == "Пользователь с таким email уже существует") || (err.Error() == "Пользователь с таким телефоном уже существует") {
			errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  http.StatusCreated,
		Data:    map[string]string{},
		Message: "OK",
	})
}

func (r resource) verifyCode(c *gin.Context) {
	var input dto.VerifyCodeRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	ctx := c.Request.Context()

	token, err := r.service.VerifyCode(ctx, input)

	if err != nil {
		if err.Error() == "user already verified" {
			errors.NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}
		if err.Error() == "user not found" || err.Error() == "invalid code" {
			errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	println(token)

	c.JSON(http.StatusOK, models.Response{
		Status: http.StatusOK,
		Data: map[string]string{
			"token": token,
		},
		Message: "OK",
	})

}

func (r resource) verifyEmail(c *gin.Context) {
	var input dto.VerifyEmailRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	err := r.service.VerifyEmail(ctx, input)

	if err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  http.StatusOK,
		Data:    map[string]string{},
		Message: "OK",
	})
}

func (r resource) forgotPassword(c *gin.Context) {
	var input dto.ForgotPasswordRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	err := r.service.ForgotPassword(ctx, input)
	if err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  http.StatusOK,
		Data:    map[string]string{},
		Message: "OK",
	})
}

func (r resource) login(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": "Not implemented",
	})
}
