package middleware

import (
	logger "../../../../infrastructure/logger/v1"
	utils "../../../../infrastructure/utils/v1"

	"gopkg.in/gin-gonic/gin.v1"
)

//GlobalTraceLogger trace all requests & responses
func GlobalTraceLogger(context *gin.Context) {
	id := utils.NewGuid()
	logger.WithField("requestId", id).Debug(context.Request)
	context.Next()
	logger.WithField("requestId", id).Debug(context.Writer)
}
