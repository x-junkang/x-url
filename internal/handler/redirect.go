package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type RedirectReq struct {
	TinyUrl string `uri:"tiny_url"`
}

func Redirect(c *gin.Context) {
	req := &RedirectReq{}
	c.ShouldBindUri(req)
	log.Info().Str("tiny_url", req.TinyUrl).Msg("receive a tiny_url request")
	// c.Redirect(http.StatusFound, "https://baidu.com")
	c.Redirect(http.StatusFound, "http://localhost:8080/ping")
}
