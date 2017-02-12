package controllers

import (
	dicProvider "../../../../dictionary/v1/business"
	httputils "../httputils"
	"github.com/gin-gonic/gin"
)

// DictionaryController implements http logic over DictionaryProvider
type DictionaryController struct {
	dicProvider dicProvider.IDictionaryPublicProvider
}

// CreateDictionaryController creates instance of dictionary controller``
func CreateDictionaryController(dbServer string, dbName string) *DictionaryController {
	var controller DictionaryController
	controller.dicProvider = dicProvider.CreateDictionaryProvider(dbServer, dbName)
	return &controller
}

//GetDictionaryItems returns all items in the dictionary
func (controller *DictionaryController) GetDictionaryItems(context *gin.Context) {
	dicCode := httputils.GetStringParameterFromPath(context, "dicCode", true)
	result, err := controller.dicProvider.GetDictionaryItems(dicCode)
	httputils.ReturnJSONResults(context, result, result == nil, err)
}

// GetDictionaryItem returns dictionary item
func (controller *DictionaryController) GetDictionaryItem(context *gin.Context) {
	dicCode := httputils.GetStringParameterFromPath(context, "dicCode", true)
	code := httputils.GetStringParameterFromPath(context, "code", true)
	result, err := controller.dicProvider.GetDictionaryItem(dicCode, code)
	httputils.ReturnJSONResults(context, result, result == nil, err)
}

// RegisterDictionaryController register routes for DictionaryController
func RegisterDictionaryController(router *gin.Engine) {
	router.GET("/dics/:dicCode", func(c *gin.Context) {
		controller := CreateDictionaryController("localhost", "dictionarydb")
		controller.GetDictionaryItems(c)
	})
	router.GET("/dics/:dicCode/:code", func(c *gin.Context) {
		controller := CreateDictionaryController("localhost", "dictionarydb")
		controller.GetDictionaryItem(c)
	})
}
