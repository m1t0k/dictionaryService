package business

import (
	Db "../db/"
	DicTypes "../types/"
)

// CreateDictionaryProvider create instance of dictionaryProvider
func CreateDictionaryProvider(dbServer string, dbName string) IDictionaryProvider {
	var provider dictionaryProvider
	provider.dbProvider = Db.CreateMongoDbDicProvider(dbServer, dbName)
	return &provider
}

/*
DictionaryProvider type
*/
type dictionaryProvider struct {
	dbProvider Db.IDicDbProvider
}

/*
insert event into mongo db
*/
func (dicProvider *dictionaryProvider) GetDictionaryList() ([]DicTypes.MetaInfoItem, error) {
	return dicProvider.dbProvider.GetDictionaryList()
}

/*
/get full event list
*/
func (dicProvider *dictionaryProvider) GetDictionaryDesc(dicCode string) (DicTypes.MetaInfoItem, error) {
	return dicProvider.dbProvider.GetDictionaryDesc(dicCode)
}

/*
get event by id
*/
func (dicProvider *dictionaryProvider) GetDictionaryItems(dicCode string) ([]DicTypes.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItems(dicCode)
}

/*
get event by id
*/
func (dicProvider *dictionaryProvider) GetDictionaryItem(dicCode string, code string) (DicTypes.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItem(dicCode, code)
}
