package server

import (
	"database/sql"
	"net/http"
	"phonebook/setup"

	"github.com/gin-gonic/gin"
)

type Address struct {
	AddressID   sql.NullInt32  `json:"addressId"`
	Description sql.NullString `json:"description"`
	City        sql.NullString `json:"city"`
	Street      sql.NullString `json:"street"`
	HomeNumber  sql.NullString `json:"homeNumber"`
	Apartment   sql.NullString `json:"apartment"`
}

type Phone struct {
	PhoneId     sql.NullInt32  `json:"phoneId"`
	Description sql.NullString `json:"description"`
	PhoneNumber sql.NullString `json:"PhoneNumber"`
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
	const query string = "SELECT * FROM contacts LEFT JOIN addresses USING (contact_id) LEFT JOIN phones USING (contact_id)"

	rows, err := db.Query(query)
	defer rows.Close()

	contacts := make(map[int]Contact)
	phones := make(map[sql.NullInt32]bool)
	addresses := make(map[sql.NullInt32]bool)

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
			if _, ok := addresses[address.AddressID]; !ok {
				val.Address = append(val.Address, address)
				addresses[address.AddressID] = true
			}
			contacts[contact.ID] = val
		} else {
			if address.AddressID.Valid == true {
				contact.Address = append(contact.Address, address)
			}
			if phone.PhoneId.Valid == true {
				contact.Phone = append(contact.Phone, phone)
			}
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

type ContactRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateContact(c *gin.Context) {
	db := setup.GetDBConn()
	var newContact ContactRequestBody

	if err := c.BindJSON(&newContact); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	const query string = "INSERT INTO contacts(first_name, last_name) VALUES (?, ?);"

	db.Exec(query, newContact.FirstName, newContact.LastName)
}
