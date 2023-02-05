package main

import "github.com/gin-gonic/gin"

func (r *repo) GetRoutes() *gin.Engine {
	router := gin.Default()
	incident := router.Group("/incident", TokenAuthentication())
	{
		incident.GET("/type/all", r.getInsidentTypes)
		incident.POST("/type", r.postInsidentType)
		incident.GET("/all", r.getIncidents)
		incident.POST("/", r.postInsident)
	}

	return router
}
