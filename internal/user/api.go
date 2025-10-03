package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandarioon/moto-alert-backend-go/internal/auth"
	"github.com/sandarioon/moto-alert-backend-go/internal/errors"
	"github.com/sandarioon/moto-alert-backend-go/internal/helpers"
	"github.com/sandarioon/moto-alert-backend-go/models/dto"
)

type resource struct {
	service Service
}

func RegisterHandlers(rg *gin.RouterGroup, service Service) {
	res := resource{service}

	rg.Use(auth.AuthMiddleware())
	rg.GET("/profile", res.getProfile)
	rg.POST("/edit", res.editUser)
	rg.POST("/updateLocation", res.updateLocation)

}

// Profile godoc
// @Summary      Profile
// @Description  Returns user data
// @Tags         user/ private
// @Produce      json
// @Success      200  {object}  dto.ProfileResponse
// @Success      401  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /user/profile [get]
func (r resource) getProfile(c *gin.Context) {
	userId, err := helpers.GetContextUserId(c)
	if err != nil {
		errors.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	ctx := c.Request.Context()

	user, err := r.service.GetProfile(ctx, nil, userId)
	if err != nil {
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.ProfileResponse{
		Status:  http.StatusOK,
		Data:    FormatUser(user),
		Message: dto.MessageOK,
	})
}

// EditUser godoc
// @Summary      Edit user
// @Description  Edit user
// @Tags         user/ private
// @Produce      json
// @Success      200  {object}  dto.ProfileResponse
// @Success      400  {object}  errors.ErrorResponse
// @Success      401  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /user/edit [post]
func (r resource) editUser(c *gin.Context) {
	userId, err := helpers.GetContextUserId(c)
	if err != nil {
		errors.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input dto.EditUserRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	user, err := r.service.EditUser(ctx, nil, userId, input)
	if err != nil {
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.ProfileResponse{
		Status:  http.StatusOK,
		Data:    FormatUser(user),
		Message: dto.MessageOK,
	})
}

// UpdateLocation godoc
// @Summary      Update location
// @Description  Update user location
// @Tags         user/ private
// @Produce      json
// @Success      200  {object}  dto.ProfileResponse
// @Success      400  {object}  errors.ErrorResponse
// @Success      401  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /user/updateLocation [post]
func (r resource) updateLocation(c *gin.Context) {
	userId, err := helpers.GetContextUserId(c)
	if err != nil {
		errors.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input dto.UpdateLocationRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	user, err := r.service.UpdateLocation(ctx, nil, userId, input)
	if err != nil {
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.ProfileResponse{
		Status:  http.StatusOK,
		Data:    FormatUser(user),
		Message: dto.MessageOK,
	})
}

// UpdateExpoPushToken godoc
// @Summary      	   Update expo push token
// @Description        Update user expo push token
// @Tags               user/ private
// @Produce            json
// @Success            200  {object}  dto.ProfileResponse
// @Success            400  {object}  errors.ErrorResponse
// @Success            401  {object}  errors.ErrorResponse
// @Failure            500  {object}  errors.ErrorResponse
// @Router       	   /user/updateExpoPushToken [post]
func (r resource) updateExpoPushToken(c *gin.Context) {
	userId, err := helpers.GetContextUserId(c)
	if err != nil {
		errors.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input dto.UpdateLocationRequest

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx := c.Request.Context()

	user, err := r.service.UpdateLocation(ctx, nil, userId, input)
	if err != nil {
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.ProfileResponse{
		Status:  http.StatusOK,
		Data:    FormatUser(user),
		Message: dto.MessageOK,
	})
}
