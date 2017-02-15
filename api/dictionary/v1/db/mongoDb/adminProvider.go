package mongoDb

import (
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
func (db *AdminDictionaryDbProvider) CreateDictionaryItem(item *types.DicItem) (bool, error) {
	if item == nil || len(item.Code) <= 0 || len(item.DCode) <= 0 {
		return false, nil // to do fix later
	}
	session, errConn := Connect(db.DbServer)
	if errConn != nil {
		return false, errConn
	}
	defer session.Close()

	var collection = GetCollection(session, db.DbName, "dics")
	err := collection.Insert(item)

	return err == nil, err
}

//UpdateDictionaryItem updates dictionary item
func (db *AdminDictionaryDbProvider) UpdateDictionaryItem(item *types.DicItem) (bool, error) {
	if item == nil || len(item.Code) <= 0 || len(item.DCode) <= 0 {
		return false, nil // to do fix later
	}

	session, errConn := Connect(db.DbServer)
	if errConn != nil {
		return false, errConn
	}
	defer session.Close()

	var collection = GetCollection(session, db.DbName, "dics")
	err := collection.Update(bson.M{"_id": item.ID}, item)

	return err == nil, err
}

//DeleteDictionaryItem deletes dictionary item
func (db *AdminDictionaryDbProvider) DeleteDictionaryItem(dicCode string, code string) (bool, error) {
	session, errConn := Connect(db.DbServer)
	if errConn != nil {
		return false, errConn
	}
	defer session.Close()

	var collection = GetCollection(session, db.DbName, "dics")
	err := collection.Remove(bson.M{"dcode": dicCode, "code": code})

	return err == nil, err
}
