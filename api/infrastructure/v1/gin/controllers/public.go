package ginControllers

import (
	"net/http"

	dicProvider "../../../../logic/v1/business/"
	"github.com/gin-gonic/gin"
)

// DictionaryController implements http logic over DictionaryProvider
type DictionaryController struct {
	dicProvider dicProvider.IDictionaryProvider
}

// GetDictionaryList returns list of dictionaries
func (controller *DictionaryController) GetDictionaryList(context *gin.Context) {
	dbProvider := dicProvider.CreateDictionaryProvider("", "")
	var result, err = dbProvider.GetDictionaryList()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, result)
	}
}

// GetDictionaryDesc returns description for the selected dictionary
func (controller *DictionaryController) GetDictionaryDesc(context *gin.Context) {
	dicCode := context.Param("dicCode")
	var result, err = controller.dicProvider.GetDictionaryDesc(dicCode)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, result)
	}
}

//GetDictionaryItems returns all items in the dictionary
func (controller *DictionaryController) GetDictionaryItems(context *gin.Context) {
	dicCode := context.Param("dicCode")
	var result, err = controller.dicProvider.GetDictionaryItems(dicCode)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, result)
	}
}

// GetDictionaryItem returns dictionary item
func (controller *DictionaryController) GetDictionaryItem(context *gin.Context) {
	dicCode := context.Param("dicCode")
	code := context.Param("code")
	var result, err = controller.dicProvider.GetDictionaryItem(dicCode, code)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, result)
	}
}

// Register DictionaryController in router
func (controller *DictionaryController) Register(router *gin.RouterGroup) {
	router.GET("/dics", controller.GetDictionaryList)
	router.GET("/dics/:dicCode", controller.GetDictionaryDesc)
	router.GET("/dics/:dicCode/:ver", controller.GetDictionaryItems)
	router.POST("/dics/:dicCode/:ver/:id", controller.GetDictionaryItem)
}
