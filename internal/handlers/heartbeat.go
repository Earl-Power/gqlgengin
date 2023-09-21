package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Heartbeat() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	}
}
