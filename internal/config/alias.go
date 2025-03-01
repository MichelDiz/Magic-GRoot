package config

import (
	"database/sql"
	"fmt"
	"log"
)

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

func UpdateAlias(newAlias, oldAlias string) {
	var projectPath string
	err := db.QueryRow(`SELECT project_path FROM aliases WHERE alias = ?`, oldAlias).Scan(&projectPath)
	if err != nil {
		log.Fatal("Erro ao buscar alias antigo:", err)
	}

	_, err = db.Exec(`INSERT INTO aliases (alias, project_path) VALUES (?, ?)`, newAlias, projectPath)
	if err != nil {
		log.Fatal("Erro ao criar novo alias:", err)
	}

	_, err = db.Exec(`DELETE FROM aliases WHERE alias = ?`, oldAlias)
	if err != nil {
		log.Fatal("Erro ao remover alias antigo:", err)
	}
}

func DeleteAlias(alias string) {
	_, err := db.Exec("DELETE FROM aliases WHERE alias = ?", alias)
	if err != nil {
		log.Fatal("Erro ao excluir alias do banco de dados:", err)
	}
}
