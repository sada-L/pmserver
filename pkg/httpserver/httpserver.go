package httpserver

import "github.com/gin-gonic/gin"

type Server struct {
	App     *gin.Engine
	Address string
}

func NewServer(port string) *Server {
	return &Server{
		App:     gin.Default(),
		Address: port,
	}
}

func (s *Server) Run() {
	addr := "localhost:" + s.Address
	s.App.Run(addr)
}
