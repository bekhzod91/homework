package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

)

func setupRouter() *gin.Engine {
	routerRoot := gin.Default()

	routerRoot.Static("/doc", "./public")
	routerRoot.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/doc")
	})

	routerV1 := routerRoot.Group("/api/v1")

	// User crud
	routerUser := routerV1.Group("/user")
	{
		userController := new(UserController)

		// List
		routerUser.GET("/", userController.list)

		// Detail
		routerUser.GET("/:id/", userController.detail)

		// Create
		routerUser.POST("/", userController.create)

		// Update
		routerUser.PUT("/:id/", userController.update)

		// Delete
		routerUser.DELETE("/:id/", userController.delete)

		// Browser cross request available
		routerUser.OPTIONS("/*any", func(c *gin.Context) {

		})
	}

	return routerRoot
}

func main() {
	initDB("data.db")
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run("0.0.0.0:8080")
}
