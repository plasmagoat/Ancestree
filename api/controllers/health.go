package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// Status godoc
// @Summary Webserver status
// @Description Returns 200 OK with "Working" when the server is running
// @Success 200 {object} string
// @Router /health [get]
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
