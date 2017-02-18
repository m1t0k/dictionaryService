package middleware

import (
	logger "../../../../infrastructure/logger/v1"
	utils "../../../../infrastructure/utils/v1"

	"github.com/gin-gonic/gin"
)

//GlobalTraceLogger trace all requests & responses
func GlobalTraceLogger(context *gin.Context) {
	id := utils.NewGuid()
	logger.WithField("requestId", id).Debug(context.Request)
	context.Next()
	logger.WithField("requestId", id).Debug(context.Writer)
}
