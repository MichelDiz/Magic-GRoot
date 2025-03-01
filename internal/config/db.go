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
		err = os.Mkdir("data", constants.DirPermissions)
		if err != nil {
			log.Fatal("Erro ao criar diretório data/:", err)
		}
	}

	var err error
	db, err = sql.Open("sqlite3", constants.DBPath)
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
		name TEXT,
		scripts TEXT
	) `)
	if err != nil {
		panic(err)
	}
}

func SetConfig(key, value string) error {
	if key == "" {
		return fmt.Errorf("config key cannot be empty")
	}

	_, err := db.Exec("INSERT INTO config (key, value) VALUES (?, ?) ON CONFLICT(key) DO UPDATE SET value = excluded.value",
		key, value)
	if err != nil {
		return fmt.Errorf("failed to set config: %w", err)
	}
	return nil
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

	rows, err := db.Query("SELECT project_path, name, scripts FROM projects")
	if err != nil {
		fmt.Println("Erro ao recuperar projetos do banco:", err)
		return projects
	}
	defer rows.Close()

	for rows.Next() {
		var projectPath string
		var name string
		var scriptsJSON string

		err := rows.Scan(&projectPath, &name, &scriptsJSON)
		if err != nil {
			fmt.Println("Erro ao escanear linha:", err)
			continue
		}

		var scripts map[string]string
		if err := json.Unmarshal([]byte(scriptsJSON), &scripts); err != nil {
			fmt.Println("Erro ao decodificar scripts JSON:", err)
			continue
		}

		scriptList := make([]string, 0, len(scripts))
		for script, command := range scripts {
			scriptList = append(scriptList, fmt.Sprintf("%s: %s", script, command))
		}

		projects[projectPath] = append([]string{name}, scriptList...)
	}

	return projects
}

func SaveProjectToDB(projectPath, name string, scripts map[string]string) {
	scriptsJSON, err := json.Marshal(scripts)
	if err != nil {
		fmt.Println("Erro ao converter scripts para JSON:", err)
		return
	}

	_, err = db.Exec(`INSERT INTO projects (project_path, name, scripts) 
		VALUES (?, ?, ?) 
		ON CONFLICT(project_path) 
		DO UPDATE SET name = excluded.name, scripts = excluded.scripts`,
		projectPath, name, string(scriptsJSON))
	if err != nil {
		fmt.Println("Erro ao salvar projeto no banco de dados:", err)
	}
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
