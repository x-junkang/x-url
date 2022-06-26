package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/x-junkang/x-url/internal/model"
)

type RedirectReq struct {
	ShotCode string `uri:"shot_code"`
}

func Redirect(c *gin.Context) {
	req := &RedirectReq{}
	c.ShouldBindUri(req)
	log.Info().Str("shot_code", req.ShotCode).Msg("receive a tiny_url request")
	retrieve, err := model.RetrieveByShotCode(req.ShotCode)
	if err != nil || retrieve == nil {
		log.Error().Err(err).Msg("retrieve tiny url fail")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// c.Redirect(http.StatusFound, "https://baidu.com")
	c.Redirect(http.StatusFound, retrieve.FullUrl)
}
