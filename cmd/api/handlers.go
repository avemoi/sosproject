package main

import (
	"context"
	"encoding/json"
	db "github.com/avemoi/sosproject/db/sqlc"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func (app *Config) getIncidents(c *gin.Context) {

	res, err := app.Models.db.ListIncidents(context.Background())
	if err != nil {
		c.String(200, "We have an error ffs")
	}
	if res == nil {
		res = make([]db.Incident, 0)
	}
	c.JSON(200, res)
}

func (app *Config) postInsident(c *gin.Context) {
	//var newInsident db.Incident
	//if err := c.BindJSON(&newInsident); err != nil {
	//	return
	//}
	//_, err := app.Models.db.CreateInsident(context.Background(), db.CreateInsidentParams{
	//	PowerID:    newInsident.PowerID,
	//	Longtitude: newInsident.Longtitude,
	//	Latitude:   newInsident.Latitude,
	//})
	//if err != nil {
	//	log.Fatal("error", err)
	//}

	c.Status(http.StatusCreated)
}

func (app *Config) UserAuthentication(c *gin.Context) {
	type powerId struct {
		Id string `json:"power_id"`
	}
	var power powerId

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}

	err = json.Unmarshal(body, &power)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	clientAddress, err := app.getAddressFromPowerID(power.Id)
	if err != nil {
		c.JSON(500, "error")
	}
	coordinates, err := app.getCoordinatesFromAddress(clientAddress)
	if err != nil {
		c.JSON(500, "error")
	}

	powerid, err := strconv.Atoi(power.Id)
	if err != nil {
		c.JSON(500, "error")
	}
	// Store to db
	userParams := db.CreateUserParams{
		PowerID:    int32(powerid),
		Latitude:   coordinates.Lat,
		Longtitude: coordinates.Lng,
	}
	res, _ := app.Models.db.CreateUser(context.Background(), userParams)
	c.JSON(200, res)
}
