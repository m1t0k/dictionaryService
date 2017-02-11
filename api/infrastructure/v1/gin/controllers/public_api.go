package ginControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m1t0k/dictionaryService/api/infrastructure/v1/config"
	middleware "github.com/m1t0k/dictionaryService/api/infrastructure/v1/gin/middleware"
	dicProvider "github.com/m1t0k/dictionaryService/api/dictionary/v1/business"
)

// DictionaryController implements http logic over DictionaryProvider
type DictionaryController struct {
	dicProvider dicProvider.IDictionaryPublicProvider
}

// CreateDictionaryController instance of DictionaryController
func CreateDictionaryController(settings config.Configuration) *DictionaryController {
	var controller DictionaryController
	controller.dicProvider = dicProvider.CreateDictionaryProvider(settings.DbServer, settings.DbName)
	return &controller
}

//GetDictionaryItems returns all items in the dictionary
func (controller *DictionaryController) GetDictionaryItems(context *gin.Context) {
	dicCode := middleware.ValidateStringParameterHanlder(context, "dicCode", true)
	var result, err = controller.dicProvider.GetDictionaryItems(dicCode)
	middleware.PassResultsToPipeLine(context, result, err)
}

// GetDictionaryItem returns dictionary item
func (controller *DictionaryController) GetDictionaryItem(context *gin.Context) {
	dicCode := middleware.ValidateStringParameterHanlder(context, "dicCode", true)
	code := middleware.ValidateStringParameterHanlder(context, "code", true)
	var result, err = controller.dicProvider.GetDictionaryItem(dicCode, code)
	middleware.PassResultsToPipeLine(context, result, err)
}

// Register DictionaryController in router
func (controller *DictionaryController) Register(router *gin.RouterGroup) {
	router.GET("/dics/:dicCode", controller.GetDictionaryItems)
	router.GET("/dics/:dicCode/:code", controller.GetDictionaryItem)
}
