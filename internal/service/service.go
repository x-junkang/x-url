package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/x-junkang/x-url/internal/config"
	"github.com/x-junkang/x-url/internal/handler"
)

type Service struct {
	listenPort string
}

func NewService(cfg *config.Config) *Service {
	return &Service{listenPort: cfg.Port}
}

func (s *Service) Run() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.POST("/tiny_url", handler.NewTinyURL)

	router.GET("/:shot_code", handler.Redirect)

	log.Info().Msg("server started!")
	router.Run(s.listenPort)
	log.Info().Msg("server stopped!")
}
