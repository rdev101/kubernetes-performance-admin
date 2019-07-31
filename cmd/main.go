package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/rdev101/kubernetes-performance-admin/internal/kube"
	"github.com/rdev101/kubernetes-performance-admin/internal/routes"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.GET("/nodes", routes.GetNodes)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
