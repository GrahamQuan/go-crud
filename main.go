package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	controllers "github.com/grahamquan/go-crud/controllers/posts"
	"github.com/grahamquan/go-crud/migrate"
	"github.com/grahamquan/go-crud/setup"
)

// runs before main()
func init() {
	setup.LoadEnv()
	setup.ConnDB()
}

func main() {
	migrate.MigrateDB()

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		v1.POST("/posts", controllers.PostsCreate)
		v1.GET("/posts", controllers.PostsGetAll)
		v1.GET("/posts/:postId", controllers.PostsGetOne)
		v1.PUT("/posts/:postId", controllers.PostsUpdate)
		v1.DELETE("/posts/:postId", controllers.PostsDelete)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Not Found"})
	})

	port := os.Getenv("PORT")
	err := r.Run(":" + port)
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
