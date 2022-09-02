package server

import (
	"net/http"
	"phonebook/setup"

	"github.com/gin-gonic/gin"
)

type Address struct {
	AddressID   int    `json:"addressId"`
	Description string `json:"description"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HomeNumber  string `json:"homeNumber"`
	Apartment   string `json:"apartment"`
}

type Phone struct {
	PhoneId     int    `json:"phoneId"`
	Description string `json:"description"`
	PhoneNumber string `json:"PhoneNumber"`
}

type Contact struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Address   []Address `json:"Address"`
	Phone     []Phone   `json:"Phone"`
}

func GetAllContacts(c *gin.Context) {
	db := setup.GetDBConn()

	rows, err := db.Query("SELECT * FROM contacts JOIN addresses USING (contact_id) JOIN phones USING (contact_id)")
	defer rows.Close()

	var contacts []Contact

	for rows.Next() {
		var contact Contact
		var address Address
		var phone Phone

		if err := rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName,
			&address.AddressID, &address.Description, &address.City, &address.Street,
			&address.HomeNumber, &address.Apartment,
			&phone.PhoneId, &phone.Description, &phone.PhoneNumber); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, contacts)
		}
		contact.Address = append(contact.Address, address)
		contact.Phone = append(contact.Phone, phone)
		contacts = append(contacts, contact)
	}

	if err = rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, contacts)
	}

	c.IndentedJSON(http.StatusOK, contacts)
}
