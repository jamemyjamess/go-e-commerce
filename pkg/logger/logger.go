package logger

import (
	"errors"
	"fmt"
	"os"
	"time"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jamemyjamess/go-e-commerce/pkg/utils"
)

type ILogger interface {
	Print() ILogger
	Save()
	SetQueryParam(c *fiber.Ctx)
	SetBodyReq(c *fiber.Ctx, iBodyReqValid IBodyReqValidation)
	SetResponse(res any)
}

type IBodyReqValidation interface {
	Validate() error
}

type bodyReqValidation struct {
	path           string
	method         string
	pathExcept     []string
	methodRequired map[string]bool
}

func NewBodyReq(path, method string) IBodyReqValidation {
	return &bodyReqValidation{
		path:   path,
		method: method,
		pathExcept: []string{
			"v1/users/signup",
		},
		methodRequired: map[string]bool{
			"POST":  true,
			"PUT":   true,
			"PATCH": true,
		},
	}
}

func (r *bodyReqValidation) Validate() error {
	// check method implemet later
	// check path
	for _, pathExcept := range r.pathExcept {
		if r.path == pathExcept {
			return errors.New("don't have permission to set body on this path")
		}
	}
	return nil
}

type Logger struct {
	Time       time.Time `json:"time"`
	TimeString string    `json:"time_string"`
	Ip         string    `json:"ip"`
	Method     string    `json:"method"`
	StatusCode int       `json:"status_code"`
	Path       string    `json:"path"`
	Query      any       `json:"query"`
	Body       any       `json:"body"`
	Response   any       `json:"response"`
}

func InitLogger(c *fiber.Ctx, res any) Logger {
	log := Logger{
		Time:       time.Now(),
		Ip:         c.IP(), // reverse proxy can't store IP
		Method:     c.Method(),
		Path:       c.Path(),
		StatusCode: c.Response().StatusCode(),
	}

	iBodyReqValid := NewBodyReq(log.Path, log.Method)
	log.SetQueryParam(c)
	log.SetBodyReq(c, iBodyReqValid)
	log.SetResponse(res)
	return log

}

func (l *Logger) setTimeString() {
	l.TimeString = l.Time.Local().Format("2006-01-02 15:04:05")
}

func (l *Logger) Print() ILogger {
	utils.PrintJsonPretty(l.Body)
	return l
}

func (l *Logger) Save() {
	filePath := fmt.Sprintf("./assets/logs/logger_%v.txt", l.Time.Format("20060102"))
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	err = errors.New("text log.Fatalf")
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer file.Close()
	data := utils.AsJson(l)
	file.WriteString(string(data) + "\n")

}

func (l *Logger) SetQueryParam(c *fiber.Ctx) {

}

func (l *Logger) SetBodyReq(c *fiber.Ctx, iBodyReqValid IBodyReqValidation) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Printf("body parser error: %v\n", err)
	}
	if err := iBodyReqValid.Validate(); err == nil {
		l.Body = body
	} else {
		log.Printf("%v\n", err)
	}

}

func (l *Logger) SetResponse(res any) {
	l.Response = res
}
