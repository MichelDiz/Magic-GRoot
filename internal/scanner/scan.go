package scanner

import (
	"encoding/json"
	"io/ioutil"
	"mgr/internal/config"
	"os"
	"path/filepath"
)

func ScanForScripts(root string) map[string][]string {
	scripts := make(map[string][]string)

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.Name() == "package.json" {
			file, _ := ioutil.ReadFile(path)
			var data map[string]interface{}
			json.Unmarshal(file, &data)
			if scriptsMap, ok := data["scripts"].(map[string]interface{}); ok {
				var scriptList []string
				for script := range scriptsMap {
					scriptList = append(scriptList, script)
				}
				scripts[path] = scriptList
				config.SaveProjectToDB(path, scriptList)
			}
		}
		return nil
	})

	return scripts
}
