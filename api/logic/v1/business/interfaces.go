package business

import (
	DicTypes "../types/"
)

// IDictionaryProvider implements read access to the dictionaries
type IDictionaryPublicProvider interface {
	GetDictionaryItems(dicCode string) ([]DicTypes.DicItem, error)
	GetDictionaryItem(dicCode string, code string) (*DicTypes.DicItem, error)
}

// IDictionaryAdminProvider implements admin access to the dictionaries
type IDictionaryAdminProvider interface {
	GetDictionaryList() ([]DicTypes.MetaInfoItem, error)
	GetDictionaryDesc(dicCode string) (*DicTypes.MetaInfoItem, error)
	CreateDictionary(dic *DicTypes.MetaInfoItem) error
	UpdateDictionary(dicCode string, dic *DicTypes.MetaInfoItem) error
	DeleteDictionary(dicCode string) error
}

// IDictionaryItemAdminProvider implements admin access to the dictionaries
type IDictionaryItemAdminProvider interface {
	GetDictionaryItems(dicCode string) ([]DicTypes.DicItem, error)
	GetDictionaryItem(dicCode string, code string) (*DicTypes.DicItem, error)
	CreateDictionaryItem(dicCode string, dicItem *DicTypes.DicItem) error
	UpdateDictionaryItem(dicCode string, dicItem *DicTypes.DicItem) error
	DeleteDictionaryItem(dicCode string, code string) error
}
