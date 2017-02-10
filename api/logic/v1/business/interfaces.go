package business

import (
	DicTypes "../types/"
)

// IDictionaryProvider implement read access to the dictionaries
type IDictionaryProvider interface {
	GetDictionaryList() ([]DicTypes.MetaInfoItem, error)
	GetDictionaryDesc(dicCode string) (*DicTypes.MetaInfoItem, error)
	GetDictionaryItems(dicCode string) ([]DicTypes.DicItem, error)
	GetDictionaryItem(dicCode string, code string) (*DicTypes.DicItem, error)
}
