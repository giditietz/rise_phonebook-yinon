package serverutils

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
}

func GetQuery(key string) (string, bool) {
	if _, ok := queryMap[key]; ok {
		return queryMap[key], ok
	}
	return "", false
}
