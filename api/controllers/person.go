package controllers

import (
	"api/httputil"
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonController struct{}

var personModel = new(models.Person)

// Get godoc
// @Summary Webserver status
// @Description Returns 200 OK with "Working" when the server is running
// @Tags person
// @Accept  json
// @Produce  json
// @Param id path int true "Person ID"
// @Success 200 {object} string
// @Router /api/person/{id} [get]
func (p PersonController) Get(c *gin.Context) {
	if c.Param("id") != "" {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(c, http.StatusInternalServerError, err, "Invalid ID")
			c.Abort()
			return
		}
		user, err := personModel.GetByID(id)
		if err != nil {
			httputil.NewError(c, http.StatusInternalServerError, err, "Error while getting person")
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Person found!", "person": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

func (p PersonController) Create(c *gin.Context) {
	if c.Param("id") != "" {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(c, http.StatusInternalServerError, err, "Invalid ID")
			c.Abort()
			return
		}
		user, err := personModel.GetByID(id)
		if err != nil {
			httputil.NewError(c, http.StatusInternalServerError, err, "Error while getting person")
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Person found!", "person": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
