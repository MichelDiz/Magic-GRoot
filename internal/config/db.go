package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_path TEXT UNIQUE,
		scripts TEXT
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

func GetProjectsFromDB() map[string][]string {
	projects := make(map[string][]string)

	rows, err := db.Query("SELECT project_path, scripts FROM projects")
	if err != nil {
		fmt.Println("Erro ao recuperar projetos do banco:", err)
		return projects
	}
	defer rows.Close()

	for rows.Next() {
		var projectPath string
		var scriptsJSON string

		err := rows.Scan(&projectPath, &scriptsJSON)
		if err != nil {
			fmt.Println("Erro ao escanear linha:", err)
			continue
		}

		var scripts []string
		json.Unmarshal([]byte(scriptsJSON), &scripts)
		projects[projectPath] = scripts
	}

	return projects
}

func SaveProjectToDB(projectPath string, scripts []string) {

	scriptsJSON, err := json.Marshal(scripts)
	if err != nil {
		fmt.Println("Erro ao converter scripts para JSON:", err)
		return
	}

	_, err = db.Exec(`INSERT INTO projects (project_path, scripts) 
		VALUES (?, ?) 
		ON CONFLICT(project_path) 
		DO UPDATE SET scripts = excluded.scripts`,
		projectPath, string(scriptsJSON))
	if err != nil {
		fmt.Println("Erro ao salvar projeto no banco de dados:", err)
	}
}

func SetAlias(alias, projectPath string) {
	_, err := db.Exec(`INSERT INTO aliases (alias, project_path) 
		VALUES (?, ?) 
		ON CONFLICT(alias) 
		DO UPDATE SET project_path = excluded.project_path`,
		alias, projectPath)
	if err != nil {
		log.Fatal("Erro ao salvar alias no banco de dados:", err)
	}
}

func GetAlias(alias string) (string, error) {
	var projectPath string
	err := db.QueryRow("SELECT project_path FROM aliases WHERE alias = ?", alias).Scan(&projectPath)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf(" Alias '%s' não encontrado.", alias)
		}
		return "", err
	}
	return projectPath, nil
}

func GetAllAliases() map[string]string {
	aliases := make(map[string]string)
	rows, err := db.Query("SELECT alias, project_path FROM aliases")
	if err != nil {
		log.Fatal("Erro ao recuperar aliases:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var alias, projectPath string
		err := rows.Scan(&alias, &projectPath)
		if err != nil {
			log.Fatal("Erro ao escanear alias:", err)
		}
		aliases[alias] = projectPath
	}

	return aliases
}

func GetScriptsFromDB(projectPath string) []string {
	var scriptsJSON string
	err := db.QueryRow("SELECT scripts FROM projects WHERE project_path = ?", projectPath).Scan(&scriptsJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(" Nenhum script encontrado para o projeto:", projectPath)
			return nil
		}
		fmt.Println("Erro ao recuperar scripts do banco de dados:", err)
		return nil
	}

	var scripts []string
	json.Unmarshal([]byte(scriptsJSON), &scripts)
	return scripts
}
