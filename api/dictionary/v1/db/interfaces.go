package db

import types "github.com/m1t0k/dictionaryService/api/dictionary/v1/types"

// IDicDbProvider implement read access to the dictionaries
type IDicDbProvider interface {
	GetDictionaryList() ([]types.MetaInfoItem, error)
	GetDictionaryDesc(dicCode string) (*types.MetaInfoItem, error)
	GetDictionaryItems(dicCode string) ([]types.DicItem, error)
	GetDictionaryItem(dicCode string, code string) (*types.DicItem, error)
}
