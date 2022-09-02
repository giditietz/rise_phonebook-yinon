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

	// var contacts []Contact
	contacts := make(map[int]Contact)
	phones := make(map[int]bool)
	addresses := make(map[int]bool)

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
		if val, ok := contacts[contact.ID]; ok {
			if _, ok := phones[phone.PhoneId]; !ok {
				val.Phone = append(val.Phone, phone)
				phones[phone.PhoneId] = true
			}
			if _, ok := phones[address.AddressID]; !ok {
				val.Address = append(val.Address, address)
				addresses[address.AddressID] = true
			}
			contacts[contact.ID] = val
		} else {
			contact.Address = append(contact.Address, address)
			contact.Phone = append(contact.Phone, phone)
			contacts[contact.ID] = contact
			phones[phone.PhoneId] = true
			addresses[address.AddressID] = true
		}
	}

	if err = rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, contacts)
	}

	c.IndentedJSON(http.StatusOK, contacts)
}
