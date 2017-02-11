package business

import (
	"log"

	db "github.com/m1t0k/dictionaryService/api/dictionary/v1/db"
	types "github.com/m1t0k/dictionaryService/api/dictionary/v1/types"
)

// CreateDictionaryProvider create instance of dictionaryProvider
func CreateDictionaryProvider(dbServer string, dbName string) IDictionaryPublicProvider {
	log.Println("=== CreateDictionaryProvider >>> ")

	var provider dictionaryProvider
	provider.dbProvider = db.CreateMongoDbDicProvider(dbServer, dbName)
	log.Println("=== CreateDictionaryProvider <<< ")

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
	log.Println("=== Provider:GetDictionaryItems >>> ")
	result, err := dicProvider.dbProvider.GetDictionaryItems(dicCode)
	log.Println("===  Provider:GetDictionaryItems <<< ")

	return result, err
}

/*
get event by id
*/
func (dicProvider *dictionaryProvider) GetDictionaryItem(dicCode string, code string) (*types.DicItem, error) {
	return dicProvider.dbProvider.GetDictionaryItem(dicCode, code)
}
