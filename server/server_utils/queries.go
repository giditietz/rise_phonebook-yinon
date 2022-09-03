package serverutils

import "fmt"

var queryMap = map[string]string{
	"getAll": `
		SELECT * FROM contacts 
		LEFT JOIN addresses USING (contact_id) 
		LEFT JOIN phones USING (contact_id)
	`,
	"insertContact": `INSERT INTO contacts (
						first_name, 
						last_name
						) 
					   VALUES (?, ?);`,
	"insertAddress": `INSERT INTO addresses(
					  contact_id, 
					  description, 
					  city, 
					  street, 
					  home_number, 
					  apartment
					  ) 
					  VALUES (?, ?, ?, ?, ?, ?)`,
	"insertPhone": `INSERT INTO phones(
					contact_id, 
					description, 
					phone_number
					) 
					VALUES (?, ?, ?)`,
	"deleteContact": "DELETE FROM contacts WHERE contact_id = ?",
	"editContact":   "UPDATE contacts SET ",
	"where":         " WHERE ",
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
