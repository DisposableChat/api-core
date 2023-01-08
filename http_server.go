package core

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Address  string
	Port     int16
	App      *fiber.App
	certFile string
	keyFile  string
}

func (s *Server) New() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	return app
}

func (s *Server) Listen() error {
	err := s.App.Listen(fmt.Sprintf("%s:%d", s.Address, s.Port))
	if err != nil {
		return err
	}

	return nil
}

func OK(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(HTTPServerMessage{
		Message: message,
		Data:    data,
		Exit:    0,
	})
}

func Error(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(HTTPServerMessage{
		Message: message,
		Data:    data,
		Exit:    1,
	})
}

func AuthorizationMiddlware(c *fiber.Ctx, required bool) (error) {
	session, err := GetSessionFromRequest(c)
	if err != nil && required {
		return Error(c, err.Error(), nil)
	}

	if err == nil && len(session.ID) > 1 {
		fmt.Println(session) //delete
		c.Locals("authorized", true)
		c.Locals("session", session)
	}

	return nil
}

func GetSessionFromRequest(c *fiber.Ctx) (Session, error) {
	var data GetSessionRequest
	var session = Session{}
	err := c.BodyParser(&data)
	if err != nil {
		return session, Error(c, BodyError, nil)
	}

	if len(data.Token) < 1 {
		return session, Error(c, InvalidTokenError, nil)
	}

	password := GetHashPassword(LocalHashPassword, data.HashPassword)
	decrypt, err := Decrypt(password, data.Token)
	if err != nil {
		return session, Error(c, InvalidTokenError, nil)
	}

	session, err = UnParseSession(decrypt)
	if err != nil {
		return session, Error(c, InvalidTokenError, nil)
	}

	return session, nil
}
