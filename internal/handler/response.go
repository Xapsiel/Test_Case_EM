package handler

import (
	"fmt"
	. "github.com/Xapsiel/EffectiveMobile/pkg/log"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message" example:"error description""`
	Status  string `json:"status" example:"fail"`
}

type resultResponse struct {
	Id     int    `json:"id,omitempty" example:"1""`
	Text   string `json:"text,omitempty" example:"description"`
	Status string `json:"status,omitempty" example:"success"`
}

func newErrorResponce(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message, Status: "fail"})
	Logger.Info(MakeLog("id должен быть числовым значением", fmt.Errorf("Status code=%d", statusCode)))
}
