package userErrorTranslator

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/jamemyjamess/go-e-commerce/pkg/errorTranslator"
)

func NewUserErrorTranslator() errorTranslator.IAppenderTranslator {
	return &UserErrorTranslator{}
}

type UserErrorTranslator struct {
	Error map[string]errorTranslator.ErrorTranslatorJSON
}

func (e *UserErrorTranslator) InitTranslator() *map[string]errorTranslator.ErrorTranslatorJSON {
	errJsonFilePath, _ := filepath.Abs("./modules/users/usersErrorTranslator/error.json")
	file, err := os.Open(errJsonFilePath)
	if err != nil {
		log.Fatalf("InitTranslator fail cause file opening error: %v", err.Error())
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&e.Error); err != nil {
		log.Fatalf("InitTranslator fail cause json decoding error: %v", err.Error())
	}
	return &e.Error
}
