package setup

import (
	"database/sql"
	"io/ioutil"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectToDB() (openError, pingError error) {

	var err error

	db, err = sql.Open("mysql", "rise:shine@tcp(db:3306)/phonebook")
	if err != nil {
		return err, nil
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return nil, nil
}

func InitSchema() error {
	path, _ := filepath.Abs("./server/setup/sql-scripts/init-schema.sql")
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDBConn() *sql.DB {
	return db
}
