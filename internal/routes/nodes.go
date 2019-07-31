package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rdev101/kubernetes-performance-admin/internal/model"
)

//GetNodes function for handling get nodes api route
func GetNodes(c *gin.Context) {
	c.JSON(200, model.Nodes{})
}
