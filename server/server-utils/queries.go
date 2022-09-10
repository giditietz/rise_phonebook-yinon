package serverutils

import "fmt"

var queryMap = map[string]string{
	"getAllContact": `SELECT * FROM contacts`,
	"insertContact": `INSERT INTO contacts (
		first_name, 
		last_name
		) 
		VALUES (?, ?);`,
	"deleteContact":     `DELETE FROM contacts WHERE contact_id = ?`,
	"editContact":       `UPDATE contacts SET `,
	"getContactAddress": `SELECT * FROM addresses WHERE contact_id = ?`,
	"insertAddress": `INSERT INTO addresses(
		contact_id, 
		description, 
		city, 
		street, 
		home_number, 
		apartment
		) 
		VALUES (?, ?, ?, ?, ?, ?)`,
	"editAddress":      `UPDATE addresses SET `,
	"deleteAddress":    `DELETE FROM addresses WHERE address_id = ? `,
	"getContactPhones": `SELECT * FROM phones WHERE contact_id = ?`,
	"insertPhone": `INSERT INTO phones(
		contact_id, 
		description, 
		phone_number
		) 
		VALUES (?, ?, ?)`,
	"editPhone":   `UPDATE phones SET `,
	"deletePhone": `DELETE FROM phones WHERE phone_id = ? `,

	"getNumOfContacts": `SELECT 
						 COUNT(contact_id)
						 FROM contacts`,
	"where": " WHERE ",
	"limit": " LIMIT",
	"and":   " AND ",
	"or":    " OR ",
	"regex": " REGEXP ",
}

const (
	keyMissingError = "no such key"
)

type KeyError struct{}

func (keyError *KeyError) Error() string {
	return keyMissingError
}

func GetQuery(key string) (string, error) {
	if _, ok := queryMap[key]; ok {
		return queryMap[key], nil
	}
	return "", &KeyError{}
}

func AddValuesToQuery(fieldName string, value string) string {
	return fmt.Sprintf(" %s = \"%s\" ", fieldName, value)
}

func RegexQuery(fieldName string, regex string) string {
	regexQuery, _ := GetQuery("regex")
	return fmt.Sprintf(" %s %s %s ", fieldName, regexQuery, regex)
}

func GetLimitQuery(offset int, limit int) (string, error) {
	ret, err := GetQuery("limit")
	if err != nil {
		return "", err
	}
	ret += fmt.Sprintf(" %d, %d", offset, limit)
	return ret, nil
}
