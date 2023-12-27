package rest

import (
	"fmt"
	"github.com/anton-kapralov/experimental/golang/game2048/internal/game"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Server struct {
	games sync.Map
}

func (s *Server) Index(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func (s *Server) NewGame(c *gin.Context) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	newGame := game.New(rng)
	key := newGameKey(rng, 16)
	s.games.Store(key, newGame)
	c.Header("Location", fmt.Sprintf("/games/%s", key))
	c.IndentedJSON(http.StatusCreated, newGame)
}

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func newGameKey(rng *rand.Rand, length int) string {
	n := len(alphabet)
	var sb strings.Builder
	for i := 0; i < length; i++ {
		idx := rng.Intn(n)
		sb.WriteRune(rune(alphabet[idx]))
	}
	return sb.String()
}

func (s *Server) GetGame(c *gin.Context) {
	key := c.Param("key")
	g, ok := s.games.Load(key)
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, g)
}

func (s *Server) MoveGame(c *gin.Context) {
	key := c.Param("key")
	g, ok := s.games.Load(key)
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}

	v := c.Query("d")
	if v != "l" && v != "r" && v != "u" && v != "d" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "d is required and must be either l, r, u, or d"})
		return
	}
	d := stringDirectionToGameDirection(v)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	newState := g.(*game.State).Move(d, rng)
	ok = s.games.CompareAndSwap(key, g, newState)
	if !ok {
		c.Status(http.StatusConflict)
	}
	c.IndentedJSON(http.StatusOK, newState)
}

func stringDirectionToGameDirection(v string) game.Direction {
	d := game.DirectionUnknown
	switch v {
	case "l":
		d = game.DirectionLeft
	case "r":
		d = game.DirectionRight
	case "u":
		d = game.DirectionUp
	case "d":
		d = game.DirectionDown
	}
	return d
}
