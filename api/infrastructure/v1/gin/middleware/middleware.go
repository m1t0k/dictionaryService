package middleware

import (
	logger "../../../../infrastructure/logger/v1"
	utils "../../../../infrastructure/utils/v1"

	"github.com/gin-gonic/gin"
)

//GlobalTraceLogger trace all requests & responses
func GlobalTraceLogger(context *gin.Context) {
	id := utils.NewGuid()
	logger.Debugf("\nRequest - %s:\n%s", id, context.Request)
	context.Next()
	logger.Debugf("\nResponse -%s:\n%s", id, context.Writer)
}
