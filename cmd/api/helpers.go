package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type externalResponse struct {
}
type clientAddress struct {
	streetName   string
	streetNumber string
	city         string
	zipCode      string
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type Result struct {
	Geometry Geometry `json:"geometry"`
}

type Response struct {
	Results []Result `json:"results"`
	Status  string   `json:"status"`
}

type Coordinates struct {
	Lat float64
	Lng float64
}

func (app *Config) getAddressFromPowerID(powerId string) (clientAddress, error) {

	var clientRes clientAddress

	url := "https://apps.deddie.gr/LicensedElectricianCertification/rest/getECOCInfoForPowerSupply"
	data := map[string]string{
		"power_supply_number": powerId,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return clientRes, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return clientRes, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return clientRes, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return clientRes, err
	}

	var result map[string]any
	err = json.Unmarshal(body, &result)
	if err != nil {
		return clientRes, err
	}

	powerSupply, ok := result["power_supply"].(map[string]interface{})
	if !ok {
		return clientRes, errors.New("not ok")
	}

	streetName, ok := powerSupply["street_name"].(string)
	if !ok {
		return clientRes, errors.New("not ok")

	}
	streetNumber, ok := powerSupply["street_number"].(string)
	if !ok {
		streetNumber = ""

	}
	zipCode, ok := powerSupply["zip_code"].(string)
	if !ok {
		return clientRes, errors.New("not ok")

	}
	city, ok := powerSupply["city"].(string)
	if !ok {
		return clientRes, errors.New("not ok")

	}

	clientRes.streetName = streetName
	clientRes.city = city
	clientRes.zipCode = zipCode
	clientRes.streetNumber = streetNumber

	return clientRes, nil

}

func (app *Config) getCoordinatesFromAddress(clientAddress clientAddress) (Coordinates, error) {
	coordinates := Coordinates{
		Lat: 0,
		Lng: 0,
	}
	address := fmt.Sprintf("%s+%s+%s+%s", clientAddress.streetName, clientAddress.streetNumber, clientAddress.zipCode, clientAddress.city)
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", address, app.GoogleApiKey)
	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// Check if the request was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return coordinates, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error:", err)
		return coordinates, err
	}
	location := response.Results[0].Geometry.Location
	coordinates.Lat = location.Lat
	coordinates.Lng = location.Lng

	return coordinates, nil
}

func (app *Config) shouldSend() (int, error) {
	// Make a GET request
	url := fmt.Sprintf("%s?time_window=%d&distance_in_m=%d", app.ClusteringAddr, app.TimeWindow, app.DistanceInMeters)
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	// Check if the request was successful
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch data: status code %d", response.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var result map[string]int

	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}

	return result["send"], nil

}
