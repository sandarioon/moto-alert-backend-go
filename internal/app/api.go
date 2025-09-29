package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandarioon/moto-alert-backend-go/internal/errors"
	"github.com/sandarioon/moto-alert-backend-go/models/dto"
)

type resource struct {
	service Service
}

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *gin.RouterGroup, service Service) {
	res := resource{service}

	r.GET("/about", res.getAbout)
	r.GET("/socialLinks", res.getSocialLinks)
	r.GET("/settings", res.getSettings)
	r.GET("/privacyPolicy", res.getPrivacyPolicy)

}

// GetAbout godoc
// @Summary      Get app description text
// @Description  Get app description text
// @Tags         app
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.GetAboutResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /app/about [get]
func (r *resource) getAbout(c *gin.Context) {
	response, err := r.service.GetAbout()
	if err != nil {
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.GetAboutResponse{
		Status: http.StatusOK,
		Data: dto.AboutData{
			Text: response,
		},
		Message: dto.MessageOK,
	})
}

// GetSocialLinks godoc
// @Summary      Get app social links
// @Description  Get app social links
// @Tags         app
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.SocialLinkResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /app/socialLinks [get]
func (r *resource) getSocialLinks(c *gin.Context) {
	socialLinks := r.service.GetSocialLinks()

	c.JSON(http.StatusOK, dto.SocialLinkResponse{
		Status:  http.StatusOK,
		Data:    socialLinks,
		Message: dto.MessageOK,
	})
}

// GetSettings godoc
// @Summary      Get app settings
// @Description  Get app settings
// @Tags         app
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.GetSettingsResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /app/settings [get]
func (r *resource) getSettings(c *gin.Context) {
	settings, err := r.service.GetSettings()
	if err != nil {
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.GetSettingsResponse{
		Status:  http.StatusOK,
		Data:    settings,
		Message: dto.MessageOK,
	})
}

// GetPrivacyPolicy godoc
// @Summary      Get privacy policy PDF
// @Description  Get privacy policy PDF
// @Tags         app
// @Accept       json
// @Produce      application/pdf
// @Success      200
// @Router       /app/privacyPolicy [get]
func (r *resource) getPrivacyPolicy(c *gin.Context) {
	filename := "privacy_policy.pdf"
	filePath := "./assets/pdf/" + filename

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/pdf")

	c.FileAttachment(filePath, filename)
}
