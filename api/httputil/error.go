package httputil

import "github.com/gin-gonic/gin"

// NewError example
func NewError(ctx *gin.Context, status int, err error, msg string) {
	er := HTTPError{
		Code:    status,
		Message: msg,
		Error:   err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
	Error   string `json:"error" example:"invalid id"`
}
