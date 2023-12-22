package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct{}

func (s *Server) Index(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
