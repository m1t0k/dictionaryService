package httputils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func respondWithError(code int, err error, context *gin.Context) {
	context.JSON(code, err.Error())
	context.Abort()
}

//GetStringParameterFromPath validates incoming params
func GetStringParameterFromPath(context *gin.Context, paramName string, required bool) string {
	value := context.Param(paramName)
	if required && len(value) <= 0 {
		respondWithError(http.StatusBadRequest, fmt.Errorf("Parameter '%s' is not a part of request", paramName), context)
	}
	return value
}

// ReturnJSONResults returns results in JSON format in case of success or returns error
func ReturnJSONResults(context *gin.Context, result interface{}, isEmptyResult bool, err error) {
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, context)
	} else if isEmptyResult {
		context.JSON(http.StatusNotFound, nil)
	} else {
		context.JSON(http.StatusOK, result)
	}
}


func Save(context *gin.Context, result interface{}, isEmptyResult bool, err error) {
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, context)
	} else if isEmptyResult {
		context.JSON(http.StatusNotFound, nil)
	} else {
		context.JSON(http.StatusOK, result)
	}
}
