package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		err = os.Mkdir("data", 0755)
		if err != nil {
			log.Fatal("Erro ao criar diretório data/:", err)
		}
	}

	var err error
	db, err = sql.Open("sqlite3", "data/magic_groot.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS config (
		key TEXT PRIMARY KEY,
		value TEXT
	)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS aliases (
		alias TEXT PRIMARY KEY,
		project_path TEXT
	)`)
	if err != nil {
		panic(err)
	}
}

func SetConfig(key, value string) {
	_, err := db.Exec("INSERT INTO config (key, value) VALUES (?, ?) ON CONFLICT(key) DO UPDATE SET value = excluded.value", key, value)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig(key string) string {
	if db == nil {
		log.Fatal("Erro: O banco de dados não foi inicializado. Certifique-se de chamar InitDB() antes de GetConfig().")
	}

	var value string
	err := db.QueryRow("SELECT value FROM config WHERE key = ?", key).Scan(&value)
	if err != nil {
		return ""
	}
	return value
}
