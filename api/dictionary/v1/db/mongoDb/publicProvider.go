package mongoDb

import (
	db "../"
	types "../../types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//PublicDictionaryDbProvider  public dictionary interface
type PublicDictionaryDbProvider struct {
	BaseMongoDbProvider
}

// CreateMongoDbDicProvider creates instance of IDicDbProvider based on mongoDbDicProvider
func CreatePublicMongoDbDictionaryProvider(dbServer string, dbName string) db.IPublicDictionaryDbProvider {
	var provider PublicDictionaryDbProvider
	provider.DbServer = dbServer
	provider.DbName = dbName
	return &provider
}

//GetDictionaryItems returns all items in the dictionary
func (db *PublicDictionaryDbProvider) GetDictionaryItems(dicCode string) ([]types.DicItem, error) {
	session, errConn := Connect(db.DbServer)
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var dicItems []types.DicItem = []types.DicItem{}
	var dicItemsCollection = GetCollection(session, db.DbName, "dics")
	err := dicItemsCollection.Find(bson.M{"dcode": dicCode}).All(&dicItems)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	return dicItems, err
}

//GetDictionaryItem returns dictionary item
func (db *PublicDictionaryDbProvider) GetDictionaryItem(dicCode string, code string) (*types.DicItem, error) {
	session, errConn := Connect(db.DbServer)
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var dicItem types.DicItem
	var dicItemsCollection = GetCollection(session, db.DbName, "dics")
	err := dicItemsCollection.Find(bson.M{"dcode": dicCode, "code": code}).One(&dicItem)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	return &dicItem, err
}
