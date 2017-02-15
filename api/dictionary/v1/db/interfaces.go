package db

import types "../types"

// IPublicDictionaryDbProvider implements public read access to the dictionaries
type IPublicDictionaryDbProvider interface {
	GetDictionaryItems(dicCode string) ([]types.DicItem, error)
	GetDictionaryItem(dicCode string, code string) (*types.DicItem, error)
}

// IAdminDictionaryDbProvider implements public read access to the dictionaries
type IAdminDictionaryDbProvider interface {
	IPublicDictionaryDbProvider
	CreateDictionaryItem(item *types.DicItem) (bool, error)
	UpdateDictionaryItem(item *types.DicItem) (bool, error)
	DeleteDictionaryItem(dicCode string, code string) (bool, error)
}

// IDicDbProvider implement read access to the dictionaries
type IDicDbProvider interface {
	GetDictionaryList() ([]types.MetaInfoItem, error)
	GetDictionaryDesc(dicCode string) (*types.MetaInfoItem, error)
	GetDictionaryItems(dicCode string) ([]types.DicItem, error)
	GetDictionaryItem(dicCode string, code string) (*types.DicItem, error)
}
