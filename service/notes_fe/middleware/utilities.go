package middleware

import (
	"atypicaldev.com/conversation/notes_fe/shared"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterRouteLogger(ctx *gin.Context) {
	logger := logrus.New()

	ctx.Set(shared.LoggerKey, logger)
	ctx.Next()
}
