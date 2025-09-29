package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandarioon/moto-alert-backend-go/internal/errors"
	"github.com/sandarioon/moto-alert-backend-go/internal/helpers"
	"github.com/sandarioon/moto-alert-backend-go/models/dto"
)

type resource struct {
	service Service
}

func RegisterHandlers(rg *gin.RouterGroup, service Service) {
	res := resource{service}

	// Public
	rg.POST("/create", res.createUser)
	rg.POST("/verifyCode", res.verifyCode)
	rg.POST("/verifyEmail", res.verifyEmail)
	rg.POST("/forgotPassword", res.forgotPassword)
	rg.POST("/resendCode", res.resendCode)
	rg.POST("/login", res.login)
	// Private
	rg.POST("/logout", AuthMiddleware(), res.logout)
}

// CreateUser godoc
// @Summary      Create user
// @Description  Creates new user if not exists
// @Tags         auth/ public
// @Accept       json
// @Produce      json
// @Param        input  body    dto.CreateUserRequest  true  "User creation request"
// @Success      200  {object}  dto.EmptyResponse
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /auth/create [post]
func (r resource) createUser(c *gin.Context) {
	var input dto.CreateUserRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	_, err := r.service.CreateUser(ctx, input)
	if err != nil {
		switch err.Error() {
		case "Пользователь с таким email уже существует", "Пользователь с таким телефоном уже существует":
			errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, dto.EmptyResponse{
		Status:  http.StatusCreated,
		Data:    dto.EmptyObject{},
		Message: dto.MessageOK,
	})
}

// VerifyCode godoc
// @Summary      Verify code
// @Description  Verify code from email
// @Tags         auth/ public
// @Accept       json
// @Produce      json
// @Param        input  body    dto.VerifyCodeRequest  true  "Verify code body"
// @Success      200  {object}  dto.VerifyCodeResponse
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      403  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /auth/verifyCode [post]
func (r resource) verifyCode(c *gin.Context) {
	var input dto.VerifyCodeRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	ctx := c.Request.Context()

	token, err := r.service.VerifyCode(ctx, input)

	if err != nil {
		switch err.Error() {
		case "user already verified":
			errors.NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		case "user not found", "invalid code":
			errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, dto.VerifyCodeResponse{
		Status: http.StatusOK,
		Data: dto.JwtToken{
			Token: token,
		},
		Message: dto.MessageOK,
	})

}

// VerifyEmail godoc
// @Summary      Verify email
// @Description  Verify if email is free
// @Tags         auth/ public
// @Accept       json
// @Produce      json
// @Param        input  body    dto.VerifyEmailRequest  true  "Verify email body"
// @Success      200  {object}  dto.EmptyResponse
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /auth/verifyEmail [post]
func (r resource) verifyEmail(c *gin.Context) {
	var input dto.VerifyEmailRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	err := r.service.VerifyEmail(ctx, input)

	if err != nil {
		switch err.Error() {
		case "user already exists":
			errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, dto.EmptyResponse{
		Status:  http.StatusOK,
		Data:    dto.EmptyObject{},
		Message: dto.MessageOK,
	})
}

// ForgotPassword godoc
// @Summary      Forgot password
// @Description  Sends email for a new password
// @Tags         auth/ public
// @Accept       json
// @Produce      json
// @Param        input  body    dto.VerifyEmailRequest  true  "Forgot password body"
// @Success      200  {object}  dto.EmptyResponse
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /auth/forgotPassword [post]
func (r resource) forgotPassword(c *gin.Context) {
	var input dto.ForgotPasswordRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	err := r.service.ForgotPassword(ctx, input)
	if err != nil {
		switch err.Error() {
		case "user not found":
			errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, dto.EmptyResponse{
		Status:  http.StatusOK,
		Data:    dto.EmptyObject{},
		Message: dto.MessageOK,
	})
}

// ResendCode godoc
// @Summary      Resend code
// @Description  Sends a new verification code
// @Tags         auth/ public
// @Accept       json
// @Produce      json
// @Param        input  body    dto.ResendCodeRequest  true  "Resend code body"
// @Success      200  {object}  dto.EmptyResponse
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /auth/resendCode [post]
func (r resource) resendCode(c *gin.Context) {
	var input dto.ResendCodeRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	err := r.service.ResendCode(ctx, input)
	if err != nil {
		switch err.Error() {
		case "user not found":
			errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, dto.EmptyResponse{
		Status:  http.StatusOK,
		Data:    dto.EmptyObject{},
		Message: dto.MessageOK,
	})
}

// Login godoc
// @Summary      Login
// @Description  login
// @Tags         auth/ public
// @Accept       json
// @Produce      json
// @Param        input  body    dto.LoginRequest  true  "Login body"
// @Success      200  {object}  dto.LoginResponse
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      403  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /auth/login [post]
func (r resource) login(c *gin.Context) {
	var input dto.LoginRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	token, err := r.service.Login(ctx, input)
	if err != nil {
		switch err.Error() {
		case "user not verified":
			errors.NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		case "user not found", "invalid password":
			errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Status: http.StatusOK,
		Data: dto.JwtToken{
			Token: token,
		},
		Message: dto.MessageOK,
	})
}

// Logout godoc
// @Summary      Logout
// @Description  Logout and remove expo push token
// @Tags         auth/ private
// @Produce      json
// @Success      200  {object}  dto.EmptyResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /auth/logout [post]
func (r resource) logout(c *gin.Context) {
	userId, err := helpers.GetContextUserId(c)
	if err != nil {
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := c.Request.Context()

	err = r.service.Logout(ctx, userId)
	if err != nil {
		switch err.Error() {
		case "user not found":
			errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, dto.EmptyResponse{
		Status:  http.StatusOK,
		Data:    dto.EmptyObject{},
		Message: dto.MessageOK,
	})
}
