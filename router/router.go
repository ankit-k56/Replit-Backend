package router

import (
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"github.com/ankit-k56/Repelit/controllers"
	"github.com/ankit-k56/Repelit/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter () *gin.Engine{
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Repelit backend",
		})
	})
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/check",middleware.CheckAuth, controllers.Validated)
	return r
	


}