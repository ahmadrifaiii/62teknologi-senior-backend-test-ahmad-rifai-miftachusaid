package ypl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Category struct {
	ID    string `json:"id"`
	Alias string `json:"alias"`
	Title string `json:"title"`
}

type Business struct {
	ID           string      `json:"id"`
	Alias        string      `json:"alias"`
	Name         string      `json:"name"`
	ImageURL     string      `json:"image_url"`
	IsClose      bool        `json:"is_close"`
	URL          string      `json:"url"`
	ReviewCount  int         `json:"review_count"`
	Categories   []Category  `json:"categories"`
	Rating       float64     `json:"rating"`
	Coordinate   Coordinates `json:"coordinates"`
	Transactions []string    `json:"transactions"`
	Price        string      `json:"price"`
	Location     Location    `json:"location"`
	Phone        string      `json:"phone"`
	DisplayPhone string      `json:"display_phone"`
	Distance     float64     `json:"distance"`
}

type Location struct {
	Address1       string   `json:"address1"`
	Address2       string   `json:"address2"`
	Address3       string   `json:"address3"`
	City           string   `json:"city"`
	State          string   `json:"state"`
	ZipCode        string   `json:"zip_code"`
	Country        string   `json:"country"`
	DisplayAddress []string `json:"display_address"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Businesses struct {
	Businesses []Business `json:"businesses"`
}

func GetData() (result *Businesses, err error) {

	var businesses *Businesses

	url := "https://api.yelp.com/v3/businesses/search?location=ID&sort_by=best_match&limit=5"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return
	}
	req.Header.Add("Authorization", "Bearer Ubf1-f0uqsJUnssqPMGo-tiFeZTT85oFmKfznlPmjDtX8s83jYMoAb-ApuD63wgq6LDZNsUXG6gurZIVYaj2jzxJmmLdCdXbDqIHU_b6KiCEVi8v-YB0OSsW6MWaY3Yx")
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &businesses)
	if err != nil {
		return
	}

	result = businesses

	return
}
