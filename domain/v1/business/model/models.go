package model

const TableBusiness = "businesses"

type Category struct {
	ID    string `json:"id"`
	Alias string `json:"alias"`
	Title string `json:"title"`
}

type Business struct {
	ID           string  `json:"id" db:"id" fieldtag:"insert"`
	Alias        string  `json:"alias" db:"alias" fieldtag:"insert,update"`
	Name         string  `json:"name" db:"name" fieldtag:"insert,update"`
	ImageURL     string  `json:"image_url" db:"image_url" fieldtag:"insert,update"`
	IsClose      bool    `json:"is_close" db:"is_close" fieldtag:"insert,update"`
	URL          string  `json:"url" db:"url" fieldtag:"insert,update"`
	ReviewCount  int     `json:"review_count" db:"review_count" fieldtag:"insert,update"`
	Rating       float64 `json:"rating" db:"rating" fieldtag:"insert,update"`
	Price        string  `json:"price" db:"price" fieldtag:"insert,update"`
	Phone        string  `json:"phone" db:"phone" fieldtag:"insert,update"`
	DisplayPhone string  `json:"display_phone" db:"display_phone" fieldtag:"insert,update"`
	Distance     float64 `json:"distance" db:"distance" fieldtag:"insert,update"`
	Deleted      bool    `json:"deleted" db:"deleted" fieldtag:"delete"`
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
