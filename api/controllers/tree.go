package controllers

import (
	"api/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TreeController ...
type TreeController struct{}

// GetFamiliyTree godoc
// @Summary Get Familiy Tree
// @Description Gets Familiy Tree for a given Person ID
// @Tags tree
// @Accept  json
// @Produce  json
// @Param id path string true "Person ID"
// @Success 200 {object} models.FamilyTree
// @Router /tree/{id} [get]
func (p TreeController) GetFamiliyTree(c *gin.Context) {
	if c.Param("id") != "" {
		tree, err := familtyTreeModel.GetForID(c.Param("id"), 4)
		if err != nil {
			httputil.NewError(c, http.StatusInternalServerError, err, "Error while getting tree")
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Done", "tree": tree})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
