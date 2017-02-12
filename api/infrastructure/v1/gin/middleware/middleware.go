package middleware

import (
	"log"

	logger "../../../../infrastructure/logger/v1"
	"github.com/gin-gonic/gin"
)

func GlobalTraceLogger(context *gin.Context) {
	logger.Debug("=== GlobalTraceLogger >>>>>>>>>")
	log.Println("=== GlobalTraceLogger >>>")
	logger.Debugf("\nRequest:\n%s", context.Request)
	context.Next()
	logger.Debugf("\nResponse:\n%s", context.Writer)
	log.Println("=== GlobalTraceLogger <<<<")

}
