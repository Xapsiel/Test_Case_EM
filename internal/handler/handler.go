package handler

import (
	"github.com/Xapsiel/EffectiveMobile/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/info", h.GetSongs)
	router.POST("/info", h.AddSong)
	router.GET("/info/verse", h.GetSongVerse)
	router.DELETE("/info", h.DeleteSong)
	router.PUT("/info", h.UpdateSong)
	return router
}
