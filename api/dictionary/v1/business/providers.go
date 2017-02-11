package business

import (
	db "github.com/m1t0k/dictionaryService/api/dictionary/v1/db"
	types "github.com/m1t0k/dictionaryService/api/logic/v1/types"
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
insert event into mongo db
*/
func (dicProvider *dictionaryProvider) GetDictionaryList() ([]types.MetaInfoItem, error) {
	return dicProvider.dbProvider.GetDictionaryList()
}

/*
/get full event list
*/
func (dicProvider *dictionaryProvider) GetDictionaryDesc(dicCode string) (*types.MetaInfoItem, error) {
	return dicProvider.dbProvider.GetDictionaryDesc(dicCode)
}

/*
get event by id
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
