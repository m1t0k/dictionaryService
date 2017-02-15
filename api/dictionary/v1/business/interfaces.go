package business

import (
	types "../types/"
)

// IPublicDictionaryProvider implements public read access to the dictionaries
type IPublicDictionaryProvider interface {
	GetDictionaryItems(dicCode string) ([]types.DicItem, error)
	GetDictionaryItem(dicCode string, code string) (*types.DicItem, error)
}

// IAdminDictionaryProvider implements admin access to the dictionaries
type IAdminDictionaryProvider interface {
	IPublicDictionaryProvider
	CreateDictionaryItem(item *types.DicItem) error
	UpdateDictionaryItem(item *types.DicItem) error
	DeleteDictionaryItem(dicCode string, code string) error
}
