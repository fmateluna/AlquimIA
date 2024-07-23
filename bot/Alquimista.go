package bot

import (
	"alquimia/config"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
)

type AlquimIABotDev struct {
	currentlyPath            string
	apiKey                   string
	messages                 []openai.ChatCompletionMessage
	response                 openai.ChatCompletionResponse
	resultPath               string
	PROMPT_REGLAS_DESARROLLO string
}

func NewIABotCTF(apiKey string) *AlquimIABotDev {

	currentlyPath, _ := os.Getwd()
	promptDesarrollador := readContentFromPath("prompt.txt")
	filenameLog := crearArchivoLog(currentlyPath)
	logFile, err := os.OpenFile(filenameLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	return &AlquimIABotDev{
		currentlyPath:            currentlyPath,
		apiKey:                   apiKey,
		PROMPT_REGLAS_DESARROLLO: promptDesarrollador,
	}
}

func (bot *AlquimIABotDev) error(mensaje string, error int) {
	log.Printf("Error %d: %s\n", error, mensaje)
}

func (bot *AlquimIABotDev) CodeMaker(config config.TaskDef) {
	resultado := config.SourcePath
	taskDefinition := config.Info
	bot.resultPath = resultado
	bot.MakeCTF(taskDefinition, resultado)
}

func (bot *AlquimIABotDev) MakeCTF(taskDefinition string, resultado string) {
	prompt := fmt.Sprintf(taskDefinition, bot.PROMPT_REGLAS_DESARROLLO)
	requestIA := prompt

	bot.messages = make([]openai.ChatCompletionMessage, 0)
	bot.messages = append(bot.messages, openai.ChatCompletionMessage{
		Role:    "assistant",
		Content: requestIA,
	})
	client := openai.NewClient(bot.apiKey)
	responseIA, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4o20240513,
			Messages: bot.messages,
		},
	)

	if err != nil {
		bot.error(fmt.Sprintf("OpenAI API returned an error: %v", err), -1)
		return
	}
	if responseIA.Choices == nil || len(responseIA.Choices) == 0 {
		bot.error("ChatGPT no ha retornado ninguna respuesta a su consulta, favor intente más tarde.", -1)
		return
	}
	bot.response = responseIA
	log.Printf("Respuesta: %s", responseIA.Choices[0].Message.Content)
	var data []ResponseIA
	err = json.Unmarshal([]byte(responseIA.Choices[0].Message.Content), &data)
	if err != nil {
		bot.error("Error decodificando la respuesta de ChatGPT", -1)
		fmt.Printf("\nRespuesta: \n%s", responseIA.Choices[0].Message.Content)
		return
	}
	for _, fileData := range data {
		fmt.Println(fileData.FileName)
		bot.createSource(
			fileData.Path,
			fileData.FileName,
			fileData.Content,
			fileData.Description,
		)
	}
}

func (bot *AlquimIABotDev) createSource(temporaryPath, filename, content, description string) {
	color.Green("Filename: %s\nDescription: %s\nContent:\n%s", filename, description, content)
	temporaryPath = filepath.FromSlash(temporaryPath)
	outputMainFolder := filepath.Join(bot.resultPath, temporaryPath)
	outputMainFolder = filepath.Dir(outputMainFolder)
	if _, err := os.Stat(outputMainFolder); os.IsNotExist(err) {
		err := os.MkdirAll(outputMainFolder, 0755)
		if err != nil {
			bot.error(fmt.Sprintf("No se pudo crear la carpeta '%s': %v", outputMainFolder, err), -1)
			return
		}
	}
	path := filepath.Join(outputMainFolder, filename)
	log.Printf("Creando archivo => %s", path)
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		bot.error(fmt.Sprintf("No se pudo escribir el archivo '%s': %v", path, err), -1)
	}
}

func (bot *AlquimIABotDev) Init(irobot string) bool {

	client := openai.NewClient(bot.apiKey)
	messages := make([]openai.ChatCompletionMessage, 0)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: irobot,
	})

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4o20240513,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	content := resp.Choices[0].Message.Content
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	bot.response = resp
	return true
}
