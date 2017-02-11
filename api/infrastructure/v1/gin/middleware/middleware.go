package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type empty struct{}

func respondWithError(code int, resp interface{}, context *gin.Context) {
	//resp := map[string]string{"error": message}
	context.JSON(code, resp)
	context.Abort()
}

//ValidateStringParameterHanlder validates incoming params
func ValidateStringParameterHanlder(context *gin.Context, paramName string, required bool) string {
	value := context.Param(paramName)
	if required && len(value) <= 0 {
		respondWithError(http.StatusBadRequest, "Parameter '"+paramName+"' is not set.", context)
	}
	return value
}

// TokenAuthMiddleware auth requests
/*
func TokenAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("api_token")

		if token == "" {
			respondWithError(401, "API token required", c)
			return
		}

		if token != os.Getenv("API_TOKEN") {
			respondWithError(401, "Invalid API token", c)
			return
		}

		c.Next()
	}
}
*/

func PassResultsToPipeLine(context *gin.Context, result interface{}, err error) {
	context.Set("error", err)
	context.Set("result", result)
}

//ProcessResultsHandler handles result from business logic
func ProcessResultsHandler(context *gin.Context) {
	err, _ := context.Get("error")
	result, _ := context.Get("result")
	if err != nil {
		respondWithError(http.StatusInternalServerError, "Error occured.", context)
	} else if result == nil {
		respondWithError(http.StatusNotFound, result, context)
	} else {
		context.JSON(http.StatusOK, result)
	}
	context.Next()
}
