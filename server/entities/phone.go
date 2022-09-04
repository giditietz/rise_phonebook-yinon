package entities

import "database/sql"

type PhoneQuery struct {
	PhoneID     sql.NullInt32  `json:"phoneID"`
	Description sql.NullString `json:"description"`
	PhoneNumber sql.NullString `json:"PhoneNumber"`
}

type PhoneRequestBody struct {
	ContactID   int    `json:"contactID"`
	PhoneID     int    `json:"phoneID"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
}

type PhoneResponseBody struct {
	PhoneID     int    `json:"PhoneID"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
}
