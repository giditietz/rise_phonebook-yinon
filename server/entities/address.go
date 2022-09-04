package entities

type AddressQuery struct {
	AddressID   int    `json:"addressID"`
	ContactID   int    `json:"contactID"`
	Description string `json:"description"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HomeNumber  string `json:"home_number"`
	Apartment   string `json:"apartment"`
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
