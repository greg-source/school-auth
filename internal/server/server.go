package server

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"school-auth/internal"
)

var _ internal.Server = &Server{}

type Server struct {
	*fiber.App
	internal.Service
}

func New(service internal.Service) *Server {
	return &Server{fiber.New(), service}
}

func (s *Server) Start(port int) error {
	s.App.Get("/profile/:username?", s.profile, s.middleware)

	return s.App.Listen(":" + strconv.Itoa(port))
}

func (s *Server) Shutdown(timeout time.Duration) error {
	return s.App.ShutdownWithTimeout(timeout)
}
