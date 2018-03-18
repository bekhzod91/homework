package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func setupRouter() *gin.Engine {
	routeRoot := gin.Default()
	routeV1 := routeRoot.Group("/api/v1")

	// User crud
	routeUser := routeV1.Group("/user")
	{
		userController := new(UserController)

		// List
		routeUser.GET("/", userController.list)

		// Detail
		routeUser.GET("/:id", userController.detail)

		// Create
		routeUser.POST("/", userController.create)

		// Update
		routeUser.PUT("/:id", userController.update)

		// Delete
		routeUser.DELETE("/:id", userController.delete)

		// Browser cross request available
		routeUser.OPTIONS("/*any", func(c *gin.Context) {

		})
	}

	return routeRoot
}

func main() {
	initDB("data.db")
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run("0.0.0.0:8080")
}
