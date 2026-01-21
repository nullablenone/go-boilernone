package routes

import "github.com/gin-gonic/gin"

func SetRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"username": "muhamad yesa",
		})
	})

	return router
}
