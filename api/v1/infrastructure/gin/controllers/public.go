package ginControllers

import (
	"net/http"

	dicProvider "/dictionaryService/api/business/providers"

	"gopkg.in/gin-gonic/gin.v1"
)

// DictionaryController implements http logic over DictionaryProvider
type DictionaryController struct {
	dicProvider dicProvider.DictionaryProvider
}

func (dc *DictionaryController) getDictionaryList(context *gin.Context) {
	var result, err = dc.dicProvider.getDictionaryList()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, result)
	}
}

func (dc *DictionaryController) getDictionaryDesc(context *gin.Context) {
	dicCode := context.Param("dicCode")
	var result, err = dc.dicProvider.getDictionaryDesc(dicCode)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, result)
	}
}

func (dc *DictionaryController) getDictionaryItems(context *gin.Context) {
	dicCode := context.Param("dicCode")
	ver := context.Param("ver")
	var result = dc.dicProvider.getDictionaryItems(dicCode, ver)
	if result == false {
		context.JSON(http.StatusInternalServerError, nil)
	} else {
		context.JSON(http.StatusOK, event)
	}
}

func (dc *DictionaryController) getDictionaryItem(context *gin.Context) {
	dicCode := context.Param("dicCode")
	ver := context.Param("ver")
	id := context.Param("id")
	var result = dc.dicProvider.getDictionaryItem(dicCode, ver, id)
	if result == false {
		context.JSON(http.StatusInternalServerError, nil)
	} else {
		context.JSON(http.StatusOK, event)
	}
}

// Register DictionaryController in router
func (dc *DictionaryController) Register(router *gin.RouterGroup) {
	router.GET("/dics", dc.getDictionaryList)
	router.GET("/dics/:dicCode", dc.getDictionaryDesc)
	router.GET("/dics/:dicCode/:ver", dc.getDictionaryItems)
	router.POST("/dics/:dicCode/:ver/:id", dc.getDictionaryItem)
}
