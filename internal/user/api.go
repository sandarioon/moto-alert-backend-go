package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type resource struct {
	service Service
}

func RegisterHandlers(r *gin.RouterGroup, service Service) {
	res := resource{service}

	r.GET("/profile", res.getProfile)

}

func (r resource) getProfile(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": "Not implemented",
	})
}
