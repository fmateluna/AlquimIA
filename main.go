package main

import (
	"alquimia/bot"
	"alquimia/config"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Ejemplo de uso
	apiKey := os.Getenv("API_KEY")
	bot := bot.NewIABotCTF(apiKey)

	jsonFilePath := flag.String("task", "task.calculadora.json", "Ruta del archivo JSON de configuración")
	flag.Parse()

	if *jsonFilePath == "" {
		fmt.Println("Uso: go run main.go -task=<ruta_del_archivo_json>")
		return
	}

	fileContent, err := ioutil.ReadFile(*jsonFilePath)
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		return
	}

	var configData config.TaskDef

	err = json.Unmarshal(fileContent, &configData)
	if err != nil {
		fmt.Printf("Error al deserializar el JSON: %v\n", err)
		return
	}

	bot.Init("Eres un desarrollador Senior que hace codigo de alto nivel. ")
	bot.CodeMaker(configData)
}
