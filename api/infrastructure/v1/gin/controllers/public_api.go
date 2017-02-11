package ginControllers

import (
	dicProvider "../../../../logic/v1/business/"
	Config "../../config/"
	Middleware "../middleware/"
	"github.com/gin-gonic/gin"
)

// DictionaryController implements http logic over DictionaryProvider
type DictionaryController struct {
	dicProvider dicProvider.IDictionaryPublicProvider
}

// CreateDictionaryController instance of DictionaryController
func CreateDictionaryController(settings Config.Configuration) *DictionaryController {
	var controller DictionaryController
	controller.dicProvider = dicProvider.CreateDictionaryProvider(settings.DbServer, settings.DbName)
	return &controller
}

//GetDictionaryItems returns all items in the dictionary
func (controller *DictionaryController) GetDictionaryItems(context *gin.Context) {
	dicCode := Middleware.ValidateStringParameterHanlder(context, "dicCode", true)
	var result, err = controller.dicProvider.GetDictionaryItems(dicCode)
	Middleware.PassResultsToPipeLine(context, result, err)
}

// GetDictionaryItem returns dictionary item
func (controller *DictionaryController) GetDictionaryItem(context *gin.Context) {
	dicCode := Middleware.ValidateStringParameterHanlder(context, "dicCode", true)
	code := Middleware.ValidateStringParameterHanlder(context, "code", true)
	var result, err = controller.dicProvider.GetDictionaryItem(dicCode, code)
	Middleware.PassResultsToPipeLine(context, result, err)
}

// Register DictionaryController in router
func (controller *DictionaryController) Register(router *gin.RouterGroup) {
	router.GET("/dics/:dicCode", controller.GetDictionaryItems)
	router.GET("/dics/:dicCode/:code", controller.GetDictionaryItem)
}
