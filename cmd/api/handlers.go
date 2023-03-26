package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	db "github.com/avemoi/sosproject/db/sqlc"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (app *Config) getIncidents(c *gin.Context) {

	timeWindow := c.Query("tw")
	if timeWindow == "" {
		res, err := app.Models.db.ListIncidents(context.Background())
		if err != nil {
			c.JSON(500, gin.H{"error": "We have an error ffs"})
			return
		}
		c.JSON(200, res)
	} else {
		res, err := app.Models.db.GetNumberOfIncidents(context.Background(), timeWindow)
		if err != nil {
			c.JSON(500, gin.H{"error": "We have an error ffs"})
			return
		}
		if res == nil {
			c.Status(http.StatusNoContent)
			return
		}
		c.JSON(200, res)
	}

}

func (app *Config) postInsident(c *gin.Context) {
	type userID struct {
		Id int `json:"user_id"`
	}
	var userId userID

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}

	err = json.Unmarshal(body, &userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Check if user has already created an incident in the time window
	userIncidentsParams := db.GetUserIncidentsParams{
		DATESUB: app.TimeWindow,
		ID:      int64(userId.Id),
	}

	userIncidents, err := app.Models.db.GetUserIncidents(context.Background(), userIncidentsParams)
	if err != nil {
		log.Fatal("error", err)
	}
	if userIncidents != nil {
		c.Status(http.StatusCreated)
		return
	}
	res, err := app.Models.db.CreateInsident(context.Background(), sql.NullInt64{
		Int64: int64(userId.Id),
		Valid: true,
	})
	if err != nil {
		log.Fatal("error", err)
	}

	fmt.Println(res)
	fmt.Println("this is a test")

	//recentIncidents, err := app.Models.db.CountRecentIncidents(context.Background(), app.TimeWindow)
	shouldSend, err := app.shouldSend()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	// If there are more than 1 incident,
	// Check clustering and if true --> send a notification
	if shouldSend == 1 { // && shouldSend() --> rest call to fastapi!!
		c.JSON(201, gin.H{
			"send": "1",
		})
		return
	}

	c.JSON(201, gin.H{
		"send": "0",
	})
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

	existingPoweridInt, err := strconv.Atoi(power.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// Check if user already exists using the power id
	existingUser, err := app.Models.db.GetUserByPowerId(context.Background(), int32(existingPoweridInt))

	// If user exists
	if err == nil && existingUser.ID != 0 {
		returnRes := make(map[string]any)
		returnRes["lat"] = existingUser.Latitude
		returnRes["lng"] = existingUser.Longitude
		returnRes["user_id"] = existingUser.ID
		c.JSON(200, returnRes)
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
		PowerID:   int32(powerid),
		Latitude:  coordinates.Lat,
		Longitude: coordinates.Lng,
	}
	res, _ := app.Models.db.CreateUser(context.Background(), userParams)
	//if err != nil {
	//	c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
	//	return
	//}
	userId, err := res.LastInsertId()
	if err != nil {
		log.Fatal("error", err)
	}
	returnRes := make(map[string]any)
	returnRes["lat"] = coordinates.Lat
	returnRes["lng"] = coordinates.Lng
	returnRes["user_id"] = userId
	c.JSON(200, returnRes)
}
