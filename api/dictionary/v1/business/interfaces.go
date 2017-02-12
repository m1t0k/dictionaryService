package business

import (
	types "../types/"
)

// IDictionaryPublicProvider implements public read access to the dictionaries
type IDictionaryPublicProvider interface {
	GetDictionaryItems(dicCode string) ([]types.DicItem, error)
	GetDictionaryItem(dicCode string, code string) (*types.DicItem, error)
}

// IDictionaryAdminProvider implements admin access to the dictionaries
type IDictionaryAdminProvider interface {
	GetDictionaryList() ([]types.MetaInfoItem, error)
	GetDictionaryDesc(dicCode string) (*types.MetaInfoItem, error)
	CreateDictionary(dic *types.MetaInfoItem) error
	UpdateDictionary(dicCode string, dic *types.MetaInfoItem) error
	DeleteDictionary(dicCode string) error
}

// IDictionaryItemAdminProvider implements admin access to the dictionaries
type IDictionaryItemAdminProvider interface {
	GetDictionaryItems(dicCode string) ([]types.DicItem, error)
	GetDictionaryItem(dicCode string, code string) (*types.DicItem, error)
	CreateDictionaryItem(dicCode string, dicItem *types.DicItem) error
	UpdateDictionaryItem(dicCode string, dicItem *types.DicItem) error
	DeleteDictionaryItem(dicCode string, code string) error
}
