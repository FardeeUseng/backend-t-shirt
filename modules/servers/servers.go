package servers

import (
	"log"

	"github.com/FardeeUseng/t-shirt-backend/configs"
	"github.com/FardeeUseng/t-shirt-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	App *fiber.App
	cfg *configs.Configs
	Db  *sqlx.DB
}

func NewServer(cfg *configs.Configs, db *sqlx.DB) *Server {
	return &Server{
		App: fiber.New(),
		cfg: cfg,
		Db:  db,
	}
}

func (s *Server) Start() {
	// if err := s.MapHandlers(); err != nil {
	// 	log.Fatalln(err.Error())
	// 	panic(err.Error())
	// }

	fiberConnURL, err := utils.ConnectionUrlBuilder("fiber", s.cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.cfg.App.Host
	port := s.cfg.App.Port
	log.Printf("server has been started on %s:%s", host, port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln("err", err.Error())
		panic(err.Error())
	}
}
