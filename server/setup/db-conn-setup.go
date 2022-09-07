package setup

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var db *sql.DB

func ConnectToDB() (openError, pingError error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   goDotEnvVariable("DBUSER"),
		Passwd: goDotEnvVariable("DBPASS"),
		Net:    "tcp",
		Addr:   goDotEnvVariable("DBURL"),
		DBName: "phonebook",
	}

	fmt.Println(cfg.FormatDSN())
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", "rise:shine@tcp(db:3306)/phonebook")
	if err != nil {
		return err, nil
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	fmt.Println("Connected!")

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
