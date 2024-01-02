package servers

import (
	"log"
	"os"
	"os/signal"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/jamemyjamess/go-e-commerce/config"
	"github.com/jamemyjamess/go-e-commerce/middlewares/middlewaresHandlers"
	"github.com/jamemyjamess/go-e-commerce/middlewares/middlewaresRepositories"
	"github.com/jamemyjamess/go-e-commerce/middlewares/middlewaresUsecases"
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

func initMiddlewares(s *Server) middlewaresHandlers.IMiddlewareHandlers {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandlers(*s.cfg, usecase)
}
func (s *Server) Start() {
	// Middlewares
	middleware := initMiddlewares(s)
	s.app.Use(middleware.CORS())

	apiV1 := s.app.Group("/api/v1")
	// cfg := *s.cfg
	modules := modules.NewModuleFactory(&apiV1, s.cfg, middleware, s.db)
	modules.MonitorModule().InitRoutes()
	modules.UserModule().InitRoutes()
	modules.UserModule().InitError()

	// Router is not founded
	s.app.Use(middleware.RouterNorFoundInfo())

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()
	log.Printf("server is starting on %v", (*s.cfg).App().URL())

	s.app.Listen((*s.cfg).App().URL())
}
