package main

import (
	"context"
	db "github.com/avemoi/sosproject/db/sqlc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (r *repo) getInsidentTypes(c *gin.Context) {

	res, err := r.db.ListIncidentTypes(context.Background())
	if err != nil {
		c.String(200, "We have an error ffs")
	}
	if res == nil {
		res = make([]db.IncidentType, 0)
	}
	c.JSON(200, res)
}

func (r *repo) postInsidentType(c *gin.Context) {
	type IncidentTypeName struct {
		Name string `json:"name"`
	}
	var incidentTypeName IncidentTypeName
	if err := c.BindJSON(&incidentTypeName); err != nil {
		return
	}
	_, err := r.db.CreateInsidentType(context.Background(), incidentTypeName.Name)
	if err != nil {
		log.Fatal("error", err)
	}
	c.Status(http.StatusCreated)

}

func (r *repo) getIncidents(c *gin.Context) {

	res, err := r.db.ListIncidents(context.Background())
	if err != nil {
		c.String(200, "We have an error ffs")
	}
	if res == nil {
		res = make([]db.Incident, 0)
	}
	c.JSON(200, res)
}

func (r *repo) postInsident(c *gin.Context) {
	var newInsident db.Incident
	if err := c.BindJSON(&newInsident); err != nil {
		return
	}
	_, err := r.db.CreateInsident(context.Background(), db.CreateInsidentParams{
		Longtitude:     newInsident.Longtitude,
		Latitude:       newInsident.Latitude,
		IncidentTypeID: newInsident.IncidentTypeID,
	})
	if err != nil {
		log.Fatal("error", err)
	}

	c.Status(http.StatusCreated)
}
