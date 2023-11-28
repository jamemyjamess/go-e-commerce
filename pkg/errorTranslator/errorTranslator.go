package errorTranslator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"
)

type IErrorTranslator interface {
	InitDefaultTranslator()
	AppendTranslator(IAppenderTranslator) error
	TranslateError(err error, language string) (string, error)
}

type IAppenderTranslator interface {
	InitTranslator() *map[string]ErrorTranslatorJSON
}

type CustomErrorTranslator struct {
	Error map[string]ErrorTranslatorJSON
}

type ErrorTranslatorJSON struct {
	Module              string            `json:"module"`
	LanguagesTranslator map[string]string `json:"languages_translator"`
}

func NewErrorTranslator() IErrorTranslator {
	e := &CustomErrorTranslator{}
	// e.InitTranslator()
	return e
}

func (c *CustomErrorTranslator) InitDefaultTranslator() {
	// errJsonPath, err := filepath.Abs("error.json")
	// if err != nil {
	// 	log.Fatalf("Get error.json's path Abs fail cause error: %v", err.Error())
	// }
	// Get the path to the error.json file
	errJsonPath := filepath.Join("pkg", "errorTranslator", "error.json")
	file, err := os.OpenFile(errJsonPath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("InitTranslator fail cause file opening error: %v", err.Error())
	}

	defer file.Close()

	if err := json.NewDecoder(file).Decode(&c.Error); err != nil {
		log.Fatalf("InitTranslator fail cause json decoding error: %v", err.Error())
	}

	// return nil
}

func (c *CustomErrorTranslator) TranslateError(err error, language string) (string, error) {
	if language == "" {
		language = "en"
	}
	_, ok := c.Error[err.Error()]
	if !ok {
		return "", fmt.Errorf("err input doesn't exist")
	}
	msg, ok := c.Error[err.Error()].LanguagesTranslator[strings.ToLower(language)]
	if !ok {
		// return "", fmt.Errorf("language not found")
		return err.Error(), nil
	}
	return msg, nil
}

func (c *CustomErrorTranslator) AppendTranslator(iAppenderTranslator IAppenderTranslator) error {

	newErrorTranslator := iAppenderTranslator.InitTranslator()
	// Merge the new error data into the existing error
	for key, value := range *newErrorTranslator {
		// validate duplicate error messages
		if _, ok := c.Error[key]; ok {
			return fmt.Errorf("some error messages is duplicated.")
		}
		c.Error[key] = value
	}
	return nil
}
