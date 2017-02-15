package mongoDb

import (
	"errors"

	db "../"
	types "../../types"
	"gopkg.in/mgo.v2/bson"
)

//AdminDictionaryDbProvider expose admin dictionary methods
type AdminDictionaryDbProvider struct {
	PublicDictionaryDbProvider
}

// CreateAdminMongoDbDictionaryProvider creates instance of IAdminDictionaryDbProvider based on mongoDbDicProvider
func CreateAdminMongoDbDictionaryProvider(dbServer string, dbName string) db.IAdminDictionaryDbProvider {
	var provider AdminDictionaryDbProvider
	provider.DbServer = dbServer
	provider.DbName = dbName
	return &provider
}

//
func (db *AdminDictionaryDbProvider) CreateDictionaryItem(item *types.DicItem) error {
	if item == nil || len(item.Code) <= 0 || len(item.DCode) <= 0 {
		return errors.New("Invalid item")
	}
	session, errConn := Connect(db.DbServer)
	if errConn != nil {
		return errConn
	}
	defer session.Close()

	var collection = GetCollection(session, db.DbName, "dics")
	return collection.Insert(item)
}

//UpdateDictionaryItem updates dictionary item
func (db *AdminDictionaryDbProvider) UpdateDictionaryItem(item *types.DicItem) error {
	if item == nil || len(item.Code) <= 0 || len(item.DCode) <= 0 {
		return errors.New("Invalid item")
	}

	session, errConn := Connect(db.DbServer)
	if errConn != nil {
		return errConn
	}
	defer session.Close()

	var collection = GetCollection(session, db.DbName, "dics")
	return collection.Update(bson.M{"_id": item.ID}, item)
}

//DeleteDictionaryItem deletes dictionary item
func (db *AdminDictionaryDbProvider) DeleteDictionaryItem(dicCode string, code string) error {
	session, errConn := Connect(db.DbServer)
	if errConn != nil {
		return errConn
	}
	defer session.Close()

	var collection = GetCollection(session, db.DbName, "dics")
	return collection.Remove(bson.M{"dcode": dicCode, "code": code})
}
