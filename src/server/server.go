package server

import (
	"log"
	"pg-conn/src/server/routes"

	"github.com/gofiber/fiber"
)

type Server struct {
	port   string
	server *fiber.App
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: fiber.New(),
	}
}

func (s *Server) Run() {
	routes := routes.ConfigRoutes(s.server)

	routes.Listen(s.port)
	log.Printf("server running on port: %s", s.port)
}
