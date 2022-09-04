package entities

type ContactQuery struct {
	ContactID    int            `json:"ContactID"`
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	AddressQuery []AddressQuery `json:"address"`
	PhoneQuery   []PhoneQuery   `json:"phone"`
}

type ContactRequestBody struct {
	FirstName  string               `json:"first_name"`
	LastName   string               `json:"last_name"`
	AddressReq []AddressRequestBody `json:"address"`
	PhoneReq   []PhoneRequestBody   `json:"phone"`
}

type ContactResponseBody struct {
	ContactID  int                   `json:"contactID"`
	FirstName  string                `json:"firstName"`
	LastName   string                `json:"lastName"`
	AddressRes []AddressResponseBody `json:"address"`
	PhoneRes   []PhoneResponseBody   `json:"phone"`
}
