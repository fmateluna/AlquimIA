package bot

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func crearArchivoLog(currentlyPath string) string {
	logPath := filepath.Join(currentlyPath, "log")
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		err := os.Mkdir(logPath, 0755)
		if err != nil {
			log.Fatalf("No se pudo crear la carpeta log: %v", err)
		}
	}
	timestamp := time.Now().Format("20060102150405")
	return filepath.Join(logPath, "log"+timestamp+".txt")
}

func readContentFromPath(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("No se pudo leer el archivo %s: %v", path, err)
	}
	return string(content)
}
