package business

import (
	db "../db/"
	mongoDb "../db/mongoDb/"
	types "../types/"
)

// CreatePublicDictionaryProvider create instance of dictionaryProvider
func CreateAdminDictionaryProvider(dbServer string, dbName string) IAdminDictionaryProvider {
	var provider AdminDictionaryProvider
	provider.dbProvider = mongoDb.CreateAdminMongoDbDictionaryProvider(dbServer, dbName)
	return &provider
}

type AdminDictionaryProvider struct {
	dbProvider db.IAdminDictionaryDbProvider
}

func (dicProvider *AdminDictionaryProvider) GetDictionaryItems(dicCode string) ([]types.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItems(dicCode)
}

func (dicProvider *AdminDictionaryProvider) GetDictionaryItem(dicCode string, code string) (*types.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItem(dicCode, code)
}

func (dicProvider *AdminDictionaryProvider) CreateDictionaryItem(item *types.DicItem) error {
	return dicProvider.dbProvider.CreateDictionaryItem(item)
}

func (dicProvider *AdminDictionaryProvider) UpdateDictionaryItem(item *types.DicItem) error {
	return dicProvider.dbProvider.UpdateDictionaryItem(item)
}

func (dicProvider *AdminDictionaryProvider) DeleteDictionaryItem(dicCode string, code string) error {
	return dicProvider.dbProvider.DeleteDictionaryItem(dicCode, code)
}
