package business

import (
	db "../db/"
	types "../types/"
)

// CreateDictionaryProvider create instance of dictionaryProvider
func CreateDictionaryProvider(dbServer string, dbName string) IDictionaryPublicProvider {
	var provider dictionaryProvider
	provider.dbProvider = db.CreateMongoDbDicProvider(dbServer, dbName)
	return &provider
}

/*
DictionaryProvider type
*/
type dictionaryProvider struct {
	dbProvider db.IDicDbProvider
}

/*
func (dicProvider *dictionaryProvider) GetDictionaryList() ([]types.MetaInfoItem, error) {
	return dicProvider.dbProvider.GetDictionaryList()
}


func (dicProvider *dictionaryProvider) GetDictionaryDesc(dicCode string) (*types.MetaInfoItem, error) {
	return dicProvider.dbProvider.GetDictionaryDesc(dicCode)
}
*/

func (dicProvider *dictionaryProvider) GetDictionaryItems(dicCode string) ([]types.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItems(dicCode)
}

/*
get event by id
*/
func (dicProvider *dictionaryProvider) GetDictionaryItem(dicCode string, code string) (*types.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItem(dicCode, code)
}
