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

	res, err := app.Models.db.CreateInsident(context.Background(), sql.NullInt64{
		Int64: int64(userId.Id),
		Valid: true,
	})
	if err != nil {
		log.Fatal("error", err)
	}

	fmt.Println(res)
	fmt.Println("this is a test")

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
