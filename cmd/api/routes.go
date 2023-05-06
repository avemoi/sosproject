package main

import "github.com/gin-gonic/gin"

func (app *Config) GetRoutes() *gin.Engine {
	router := gin.Default()
	authentication := router.Group("/auth", TokenAuthentication())
	{
		authentication.POST("/user", app.UserAuthentication)
	}
	incident := router.Group("/incident", TokenAuthentication())
	{
		incident.GET("/all", app.getIncidents)
		incident.POST("", app.postInsident)
	}

	return router
}
