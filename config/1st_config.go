package config

import (
	"flag"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var appMode string

func envPath() string {
	// filepath.Abs if path is not absolute then return working directory join with filename input
	// if len(os.Args) == 1 {
	// 	return filepath.Abs("./.env")
	// } else {
	// 	return filepath.Abs("./" + os.Args[1])
	// }

	// application mode setting to use a environment file
	flag.StringVar(&appMode, "mode", "dev", "Set the application mode (dev, prod)")
	flag.Parse()

	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Couldn't get working directory err: %v", err)
	}
	if len(os.Args) == 1 {
		return currentWorkingDirectory + "/.env." + appMode
	} else {
		return currentWorkingDirectory + os.Args[1]
	}

}

func InitConfig() IConfig {
	// filename and line of code for log package
	log.SetFlags(log.Lshortfile | log.LstdFlags)

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
		db: &db{
			host: envMap["POSTGRES_DB_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["POSTGRES_DB_PORT"])
				if err != nil {
					log.Fatalf("load db port failed: %v", err)
				}
				return p
			}(),
			protocol: envMap["POSTGRES_DB_PROTOCOL"],
			username: envMap["POSTGRES_DB_USERNAME"],
			password: envMap["POSTGRES_DB_PASSWORD"],
			database: envMap["POSTGRES_DB_DATABASE"],
			sslMode:  envMap["POSTGRES_DB_SSL_MODE"],
			maxConnections: func() int {
				m, err := strconv.Atoi(envMap["POSTGRES_DB_MAX_CONNECTIONS"])
				if err != nil {
					log.Fatalf("load db max connections failed: %v", err)
				}
				return m
			}(),
		},
		jwt: &jwt{
			adminKey:  envMap["JWT_ADMIN_KEY"],
			secretKey: envMap["JWT_SECRET_KEY"],
			apiKey:    envMap["JWT_API_KEY"],
			accessExpiresAt: func() int {
				t, err := strconv.Atoi(envMap["JWT_ACCESS_EXPIRES"])
				if err != nil {
					log.Fatalf("load access expires at failed: %v", err)
				}
				return t
			}(),
			refreshExpiresAt: func() int {
				t, err := strconv.Atoi(envMap["JWT_REFRESH_EXPIRES"])
				if err != nil {
					log.Fatalf("load refresh expires at failed: %v", err)
				}
				return t
			}(),
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
