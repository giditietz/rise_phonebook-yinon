package entities

type PhoneQuery struct {
	PhoneID     int    `json:"phoneID"`
	ContactID   int    `json:"contactID"`
	Description string `json:"description"`
	PhoneNumber string `json:"PhoneNumber"`
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
