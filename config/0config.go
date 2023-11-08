package config

import (
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func envPath() string {
	// filepath.Abs if path is not absolute then return working directory join with filename input
	// if len(os.Args) == 1 {
	// 	return filepath.Abs("./.env")
	// } else {
	// 	return filepath.Abs("./" + os.Args[1])
	// }
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Couldn't get working directory err: %v", err)
	}
	if len(os.Args) == 1 {
		return currentWorkingDirectory + "/.env"
	} else {
		return currentWorkingDirectory + os.Args[1]
	}

}

func InitConfig() IConfig {
	envMap, err := godotenv.Read(envPath())
	if err != nil {
		log.Fatalf("Couldn't read environment err: %v", err)
	}

	stringToInt := func(str string) int {
		v, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Couldn't parse APP_PORT string to int err: %v", err)
		}
		return v
	}
	secondStringToTimeDuration := func(t string) time.Duration {
		v, err := strconv.Atoi(t)
		if err != nil {
			log.Fatalf("Couldn't parse string to int err: %v", err)
		}
		return time.Duration(int64(v) * int64(math.Pow10(9)))
	}

	return &config{
		app: &app{
			host:         envMap["APP_HOST"],
			port:         stringToInt(envMap["APP_PORT"]),
			name:         envMap["APP_NAME"],
			version:      envMap["APP_VERSION"],
			readTimeout:  secondStringToTimeDuration(envMap["APP_READ_TIMEOUT"]),
			writeTimeout: secondStringToTimeDuration(envMap["APP_WRITE_TIMEOUT"]),
			bodyLimit:    stringToInt(envMap["APP_BODY_LIMIT"]),
			fileLimit:    stringToInt(envMap["APP_FILE_LIMIT"]),
			gcpBucket:    envMap[""],
		},
	}

}

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

func (c *config) App() IAppConfig {
	return c.app
}

func (c *config) Db() IDbConfig {
	return c.db
}
func (c *config) Jwt() IJwtConfig {
	return c.jwt
}
