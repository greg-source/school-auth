package server

import (
	"github.com/gofiber/fiber/v3"
)

func (s *Server) middleware(c fiber.Ctx) error {
	apiKey := c.Get("Api-key")

	ok, err := s.Service.ValidateAPIKey(c.Context(), apiKey)
	if err != nil {
		return c.SendStatus(400)
	}

	if !ok {
		return c.SendStatus(403)
	}

	return c.Next()
}
