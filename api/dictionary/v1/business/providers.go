package business

import (
	db "../db/"
	mongoDb "../db/mongoDb/"
	types "../types/"
)

// CreateDictionaryProvider create instance of dictionaryProvider
func CreateDictionaryProvider(dbServer string, dbName string) IPublicDictionaryProvider {
	var provider dictionaryProvider
	provider.dbProvider = mongoDb.CreatePublicMongoDbDictionaryProvider(dbServer, dbName)
	return &provider
}

/*
DictionaryProvider type
*/
type dictionaryProvider struct {
	dbProvider db.IPublicDictionaryDbProvider
}

func (dicProvider *dictionaryProvider) GetDictionaryItems(dicCode string) ([]types.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItems(dicCode)
}

func (dicProvider *dictionaryProvider) GetDictionaryItem(dicCode string, code string) (*types.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItem(dicCode, code)
}
