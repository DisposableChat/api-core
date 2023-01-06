package core

import (
	"fmt"

	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

type Server struct {
	Address    string
	Port       int16
	HTTPServer fiber.App
	Listener   quic.Listener
	certFile   string
	keyFile    string
}

func (s *Server) New() fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	return *app
}

func (s *Server) Listen() error {
	httpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.HTTPServer.Handler()
	})

	err := http3.ListenAndServe(fmt.Sprintf("%s:%d", s.Address, s.Port), s.certFile, s.keyFile, httpHandler)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
