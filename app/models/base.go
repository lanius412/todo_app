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

/*
const (
	tableNameUser = "users"
	tableNameTodo = "todos"
	tableNameSession = "sessions"
)
*/

func init() {
	//postgresql ver
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}

	/* mysql ver
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatal(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				uuid STRING NOT NULL UNIQUE,
				name STRING,
				email STRING,
				password STRING,
				created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				content TEXT,
				user_id INTEGER,
				created_at DATETIME)`, tableNameTodo)
	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				uuid STRING NOT NULL UNIQUE,
				email STRING,
				user_id INTEGER,
				created_at DATETIME)`, tableNameSession)
	Db.Exec(cmdS)
	*/
}

func createUUID() (uuidObj uuid.UUID) {
	uuidObj, _ = uuid.NewUUID()
	return uuidObj
}

func Encrypt(plaintext string) (crypttext string) {
	crypttext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return crypttext
}