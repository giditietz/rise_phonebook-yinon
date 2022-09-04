package entities

import "database/sql"

type AddressQuery struct {
	AddressID   sql.NullInt32  `json:"addressID"`
	Description sql.NullString `json:"description"`
	City        sql.NullString `json:"city"`
	Street      sql.NullString `json:"street"`
	HomeNumber  sql.NullString `json:"home_number"`
	Apartment   sql.NullString `json:"apartment"`
}

type AddressRequestBody struct {
	ContactID   int    `json:"contactID"`
	AddressID   int    `json:"addressID"`
	Description string `json:"description"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HomeNumber  string `json:"home_number"`
	Apartment   string `json:"apartment"`
}

type AddressResponseBody struct {
	AddressID   int    `json:"AddressID"`
	Description string `json:"description"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HomeNumber  string `json:"home_number"`
	Apartment   string `json:"apartment"`
}
