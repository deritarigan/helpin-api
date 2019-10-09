package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/subosito/gotenv"

	"github.com/lib/pq"
	// _ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	users    = "deritarigan"
	password = "deri123"
	dbname   = "helpin"
)

var db *sql.DB
var err error

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gotenv.Load()
}

//ConnectDB to Online db
func ConnectDB() *sql.DB {
	pqURL, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	logFatal(err)

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, users, password, dbname)

	db, err = sql.Open("postgres", pqURL)
	// logFatal(err)
	err = db.Ping()
	// logFatal(err)

	return db
}

//Connect to Local db
func ConnectDBLocal() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, users, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	logFatal(err)
	err = db.Ping()

	return db
}
