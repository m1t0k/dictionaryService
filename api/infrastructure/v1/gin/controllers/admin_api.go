package controllers

import (
	dicProvider "../../../../dictionary/v1/business"
	types "../../../../dictionary/v1/types"
	config "../../config"
	httputils "../httputils"
	"github.com/gin-gonic/gin"
)

// AdminDictionaryController implements http logic over DictionaryProvider
type AdminDictionaryController struct {
	dicProvider dicProvider.IAdminDictionaryProvider
}

// CreateDictionaryAdminController instance of DictionaryController
func CreateAdminDictionaryController(dbServer string, dbName string) *AdminDictionaryController {
	var controller AdminDictionaryController
	controller.dicProvider = dicProvider.CreateAdminDictionaryProvider(dbServer, dbName)
	return &controller
}

//GetDictionaryItems returns all items in the dictionary
func (controller *AdminDictionaryController) GetDictionaryItems(context *gin.Context) {
	dicCode := httputils.GetStringParameterFromPath(context, "dicCode", true)
	result, err := controller.dicProvider.GetDictionaryItems(dicCode)
	httputils.ReturnJSONResults(context, result, result == nil, err)
}

// GetDictionaryItem returns dictionary item
func (controller *AdminDictionaryController) GetDictionaryItem(context *gin.Context) {
	dicCode := httputils.GetStringParameterFromPath(context, "dicCode", true)
	code := httputils.GetStringParameterFromPath(context, "code", true)
	result, err := controller.dicProvider.GetDictionaryItem(dicCode, code)
	httputils.ReturnJSONResults(context, result, result == nil, err)
}

//GetDictionaryItems returns all items in the dictionary
func (controller *AdminDictionaryController) CreateDictionaryItem(context *gin.Context) {
	var item types.DicItem
	context.BindJSON(&item)
	err := controller.dicProvider.CreateDictionaryItem(&item)
	httputils.ReturnJSONResults(context, err == nil, err == nil, err)
}

// GetDictionaryItem returns dictionary item
func (controller *AdminDictionaryController) UpdateDictionaryItem(context *gin.Context) {
	//dicCode := httputils.GetStringParameterFromPath(context, "dicCode", true)
	var item types.DicItem
	context.BindJSON(&item)
	err := controller.dicProvider.UpdateDictionaryItem(&item)
	httputils.ReturnJSONResults(context, err == nil, err == nil, err)
}

// DeleteDictionaryItem deletes dictionary item
func (controller *AdminDictionaryController) DeleteDictionaryItem(context *gin.Context) {
	dicCode := httputils.GetStringParameterFromPath(context, "dicCode", true)
	code := httputils.GetStringParameterFromPath(context, "code", true)
	err := controller.dicProvider.DeleteDictionaryItem(dicCode, code)
	httputils.ReturnJSONResults(context, err == nil, err == nil, err)
}

// RegisterAdminDictionaryController registers routes for DictionaryAdminController
func RegisterAdminDictionaryController(router *gin.RouterGroup, configSettings config.Configuration) {
	router.GET("/dics/:dicCode", func(c *gin.Context) {
		controller := CreateAdminDictionaryController(configSettings.DbServer, configSettings.DbName)
		controller.GetDictionaryItems(c)
	})

	router.GET("/dics/:dicCode/:code", func(c *gin.Context) {
		controller := CreateAdminDictionaryController(configSettings.DbServer, configSettings.DbName)
		controller.GetDictionaryItem(c)
	})

	router.POST("/dics", func(c *gin.Context) {
		controller := CreateAdminDictionaryController(configSettings.DbServer, configSettings.DbName)
		controller.CreateDictionaryItem(c)
	})

	router.PUT("/dics/:dicCode", func(c *gin.Context) {
		controller := CreateAdminDictionaryController(configSettings.DbServer, configSettings.DbName)
		controller.UpdateDictionaryItem(c)
	})

	router.DELETE("/dics/:dicCode/:code", func(c *gin.Context) {
		controller := CreateAdminDictionaryController(configSettings.DbServer, configSettings.DbName)
		controller.DeleteDictionaryItem(c)
	})
}
