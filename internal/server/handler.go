package server

import (
	"github.com/gofiber/fiber/v3"
)

func (s *Server) profile(c fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		profiles, err := s.Service.Get(c.Context())
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(profiles)
	} else {
		profile, err := s.Service.GetByUsername(c.Context(), username)
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(profile)
	}
}
