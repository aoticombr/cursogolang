package main

import (
	"fmt"
	"time"
)

type Server struct {
	port           int
	timeout        time.Duration
	maxConnections int
}
type ServerOptions func(*Server)

func SetPort(value int) ServerOptions {
	return func(s *Server) {
		s.port = value
	}
}
func SetTimeOut(value time.Duration) ServerOptions {
	return func(s *Server) {
		s.timeout = value
	}
}
func SetMaxConn(value int) ServerOptions {
	return func(s *Server) {
		s.maxConnections = value
	}
}

/*
Modo Antigo
func NewServer() *Server {
	return &Server{
		port:           8080,
		timeout:        time.Second * 10,
		maxConnections: 100,
	}
}*/

func NewServer(options ...ServerOptions) *Server {
	server := &Server{
		port:           8080,
		timeout:        time.Second * 10,
		maxConnections: 100,
	}

	for _, option := range options {
		option(server)
	}

	return server
}

func main() {
	fmt.Println("Aula 25 - Functional Options Pattern(FOP)/Padrão de Opções Funcionais")

	Server := NewServer(
		SetPort(9001),
		SetMaxConn(700),
	)

	fmt.Printf("Config do Server? %+v\n", Server)
}
