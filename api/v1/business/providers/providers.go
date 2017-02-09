package providers

import (
	MongoDb "../db/mongo/"
	DicTypes "../types/"
)

/*
DictionaryProvider type
*/
type DictionaryProvider struct {
	dbProvider MongoDb.MongoDbDicProvider
}

/*
insert event into mongo db
*/
func (dicProvider *DictionaryProvider) GetDictionaryList() ([]DicTypes.MetaInfoItem, error) {
	return dicProvider.dbProvider.GetDictionaryList()
}

/*
/get full event list
*/
func (dicProvider *DictionaryProvider) GetDictionaryDesc(dicCode string) (DicTypes.MetaInfoItem, error) {
	return dicProvider.dbProvider.GetDictionaryDesc(dicCode)
}

/*
get event by id
*/
func (dicProvider *DictionaryProvider) GetDictionaryItems(dicCode string) ([]DicTypes.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItems(dicCode)
}

/*
get event by id
*/
func (dicProvider *DictionaryProvider) GetDictionaryItem(dicCode string, code string) (DicTypes.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItem(dicCode, code)
}
