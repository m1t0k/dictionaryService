package business

import (
	db "../db"
	mongoDb "../db/mongoDb"
	types "../types"
)

// CreatePublicDictionaryProvider create instance of dictionaryProvider
func CreatePublicDictionaryProvider(dbServer string, dbName string) IPublicDictionaryProvider {
	var provider PublicDictionaryProvider
	provider.dbProvider = mongoDb.CreatePublicMongoDbDictionaryProvider(dbServer, dbName)
	return &provider
}

/*
DictionaryProvider type
*/
type PublicDictionaryProvider struct {
	dbProvider db.IPublicDictionaryDbProvider
}

func (dicProvider *PublicDictionaryProvider) GetDictionaryItems(dicCode string) ([]types.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItems(dicCode)
}

func (dicProvider *PublicDictionaryProvider) GetDictionaryItem(dicCode string, code string) (*types.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItem(dicCode, code)
}
