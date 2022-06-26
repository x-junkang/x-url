package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/x-junkang/x-url/internal/model"
	"github.com/x-junkang/x-url/pkg/utils"
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
	//TODO verify
	shot, err := utils.GenerateID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = model.CreateTinyUrl(req.Url, shot)
	if err != nil {
		log.Error().Err(err).Msg("create tiny url fail")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": req.Url, "tiny_url": shot})
}
