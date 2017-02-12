package controllers

import (
	dicProvider "../../../../dictionary/v1/business"
	config "../../config"
	"github.com/gin-gonic/gin"
)

// DictionaryAdminController implements http logic over DictionaryProvider
type DictionaryAdminController struct {
	dicProvider dicProvider.IDictionaryAdminProvider
}

// CreateDictionaryAdminController instance of DictionaryController
func CreateDictionaryAdminController(settings config.Configuration) *DictionaryAdminController {
	var controller DictionaryAdminController
	//controller.dicProvider = dicProvider.CreateDictionaryAdminProvider(settings.DbServer, settings.DbName)
	return &controller
}

// GetDictionaryList returns list of dictionaries
func (controller *DictionaryAdminController) GetDictionaryList(context *gin.Context) {
	//var result, err = controller.dicProvider.GetDictionaryList()
	//Middleware.ProcessResultsHandler(result, err, context)
}

// GetDictionaryDesc returns description for the selected dictionary
func (controller *DictionaryAdminController) GetDictionaryDesc(context *gin.Context) {
	//dicCode := Middleware.ValidateStringParameterHanlder(context, "dicCode", true)
	//	var result, err = controller.dicProvider.GetDictionaryDesc(dicCode)
	//	Middleware.ProcessResultsHandler(result, err, context)
}

// Register DictionaryController in router
/*func (controller *DictionaryAdminController) Register(router *gin.RouterGroup) {
	router.GET("/meta", controller.GetDictionaryList)
	router.GET("/meta/:dicCode", controller.GetDictionaryDesc)
}*/
