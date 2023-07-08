package main

import "fmt"

type ForeignUser struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Address  ForeignAddress `json:"address"`
	Phone    string         `json:"phone"`
	Website  string         `json:"website"`
	Company  ForeignCompany `json:"company"`
}

type ForeignAddress struct {
	Street  string     `json:"street"`
	Suite   string     `json:"suite"`
	City    string     `json:"city"`
	Zipcode string     `json:"zipcode"`
	Geo     ForeignGeo `json:"geo"`
}

func (a ForeignAddress) FormatAddr() string {
	return fmt.Sprintf("%s %s (%s, %s)", a.City, a.Zipcode, a.Geo.Lat, a.Geo.Lng)
}

type ForeignGeo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type ForeignCompany struct {
	Name string `json:"name"`
	// Omitted because of curiosity
	// CatchPhrase string `json:"catchPhrase"`
	Bs string `json:"bs"`
}
