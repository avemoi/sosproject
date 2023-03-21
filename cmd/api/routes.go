package main

import "github.com/gin-gonic/gin"

func (app *Config) GetRoutes() *gin.Engine {
	router := gin.Default()
	incident := router.Group("/incident", TokenAuthentication())
	{
		incident.GET("/all", app.getIncidents)
		incident.POST("/", app.postInsident)
		incident.POST("/authentication", app.UserAuthentication)
	}

	return router
}
