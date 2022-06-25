package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TinyURLReq struct {
	Url string `json:"url"`
}

func NewTinyURL(c *gin.Context) {
	req := &TinyURLReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info().Str("url", req.Url).Msg("receive a request")
	c.JSON(http.StatusOK, gin.H{"url": req.Url, "tiny_url": "test"})
}
