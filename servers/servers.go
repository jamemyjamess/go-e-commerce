package servers

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/jamemyjamess/go-e-commerce/config"
	"github.com/jamemyjamess/go-e-commerce/modules"

	"github.com/jmoiron/sqlx"
)

type IServer interface {
	Start()
}

type Server struct {
	app *fiber.App
	cfg *config.IConfig
	db  *sqlx.DB
}

func NewServer(cfg *config.IConfig, db *sqlx.DB) IServer {
	return &Server{
		cfg: cfg,
		db:  db,
		app: fiber.New(fiber.Config{
			AppName:     (*cfg).App().Name(),
			BodyLimit:   (*cfg).App().BodyLimit(),
			ReadTimeout: (*cfg).App().ReadTimeOut(),
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
}

func (s *Server) Start() {
	apiV1 := s.app.Group("/api/v1")
	cfgApp := (*s.cfg).App()
	modules.NewModuleFactory(&apiV1, &cfgApp)
}
