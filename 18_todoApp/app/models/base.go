package models

// TBL DBの作成
import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todoAPP/config"
	"github.com/google/uuid"
	_	"github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

const (
	tableNameUser ="users"
)

func init() {
	DB, err = sql.Open(config.Config.SQLDriver,config.Config.DBName)

	if err !=nil {
		log.Fatalln(err)
	}
	cmdU := fmt.Sprintf(
	`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME
		)`, tableNameUser)
	DB.Exec(cmdU)
}

func createUUID() (uuidObj uuid.UUID) {
	uuidObj, _ = uuid.NewUUID()
	return uuidObj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}