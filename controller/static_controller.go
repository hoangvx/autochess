package controller

import (
	"github.com/gin-gonic/gin"
)

// Controller is user controlller
type StaticController struct{}

// Index action: GET /users
func (pc StaticController) Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}