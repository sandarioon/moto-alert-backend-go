package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandarioon/moto-alert-backend-go/internal/errors"
	"github.com/sandarioon/moto-alert-backend-go/models"
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
// @Success      200  {object}  models.Response
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /app/about [get]
func (r *resource) getAbout(c *gin.Context) {
	response, err := r.service.GetAbout()
	if err != nil {
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status: http.StatusOK,
		Data: map[string]string{
			"text": response,
		},
		Message: "OK",
	})
}

func (r *resource) getSocialLinks(c *gin.Context) {
	socialLinks, err := r.service.GetSocialLinks()
	if err != nil {
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  http.StatusOK,
		Data:    socialLinks,
		Message: "OK",
	})
}

func (r *resource) getSettings(c *gin.Context) {
	settings, err := r.service.GetSettings()
	if err != nil {
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  http.StatusOK,
		Data:    settings,
		Message: "OK",
	})
}

func (r *resource) getPrivacyPolicy(c *gin.Context) {
	filename := "privacy_policy.pdf"
	filePath := "./assets/pdf/" + filename

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/pdf")

	c.FileAttachment(filePath, filename)
}
