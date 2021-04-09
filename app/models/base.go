package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"

	"todo_app/config"

	"github.com/lib/pq"
	"github.com/google/uuid"
)

var Db *sql.DB
var err error

func init() {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}
}

func createUUID() (uuidObj uuid.UUID) {
	uuidObj, _ = uuid.NewUUID()
	return uuidObj
}

func Encrypt(plaintext string) (crypttext string) {
	crypttext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return crypttext
}