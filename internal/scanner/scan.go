package scanner

import (
	"encoding/json"
	"mgr/internal/config"
	"os"
	"path/filepath"
	"strings"
)

type PackageJSON struct {
	Name    string            `json:"name"`
	Scripts map[string]string `json:"scripts"`
}

func ScanForScripts(rootPath string) map[string]map[string]string {
	projects := make(map[string]map[string]string)
	_ = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && strings.HasSuffix(path, "node_modules") {
			return filepath.SkipDir
		}

		if info.Name() == "package.json" {
			packageData, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var pkg PackageJSON
			if err := json.Unmarshal(packageData, &pkg); err != nil {
				return err
			}

			projectPath := filepath.Dir(path)
			projects[projectPath] = pkg.Scripts

			config.SaveProjectToDB(projectPath, pkg.Name, pkg.Scripts)
		}
		return nil
	})
	return projects
}
