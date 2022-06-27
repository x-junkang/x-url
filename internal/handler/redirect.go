package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/x-junkang/x-url/internal/cache"
	"github.com/x-junkang/x-url/internal/model"
)

type RedirectReq struct {
	ShotCode string `uri:"shot_code"`
}

func Redirect(c *gin.Context) {
	req := &RedirectReq{}
	err := c.ShouldBindUri(req)
	if err != nil {
		log.Error().Err(err).Msg("parse request fail")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shot := req.ShotCode
	log.Info().Str("shot_code", shot).Msg("receive a tiny_url request")

	//read cache
	tryCache, err := cache.GetKey(c, shot)
	if err != nil {
		log.Err(err).Msg("read cache failed")
	}
	if tryCache != "" {
		log.Debug().Str("url", tryCache).Msg("retrieve shot code by cache")
		c.Redirect(http.StatusFound, tryCache)
		return
	}

	//retrieve from database
	retrieve, err := model.RetrieveByShotCode(shot)
	if err != nil || retrieve == nil {
		log.Error().Err(err).Msg("retrieve tiny url fail")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//set cache
	err = cache.SetKeyValue(c, shot, retrieve.FullUrl)
	if err != nil {
		log.Err(err).Msg("set cache failed")
	}
	c.Redirect(http.StatusFound, retrieve.FullUrl)
}
