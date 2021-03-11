package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController ...
type HealthController struct{}

// Status godoc
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
