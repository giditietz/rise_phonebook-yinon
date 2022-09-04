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
	"getContactPhones": `SELECT * FROM phones WHERE contact_id = ?`,
	"insertPhone": `INSERT INTO phones(
					contact_id, 
					description, 
					phone_number
					) 
					VALUES (?, ?, ?)`,
	"editPhone": `UPDATE phones SET `,
	"where":     " WHERE ",
	"limit":     " LIMIT",
	"and":       " AND ",
}

func GetQuery(key string) (string, bool) {
	if _, ok := queryMap[key]; ok {
		return queryMap[key], ok
	}
	return "", false
}

func AddValuesToQuery(fieldName string, value string) string {
	return fmt.Sprintf(" %s = \"%s\" ", fieldName, value)
}

func GetLimitQuery(offset int, limit int) string {
	ret, _ := GetQuery("limit")
	ret += fmt.Sprintf(" %d, %d", offset, limit)
	return ret
}
