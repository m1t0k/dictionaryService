package db

import (
	"log"

	types "../types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreateMongoDbDicProvider creates instance of IDicDbProvider based on mongoDbDicProvider
func CreateMongoDbDicProvider(dbServer string, dbName string) IDicDbProvider {
	var provider mongoDbDicProvider
	provider.DbServer = dbServer
	provider.DbName = dbName
	return &provider
}

type mongoDbDicProvider struct {
	DbServer string
	DbName   string
}

func (db *mongoDbDicProvider) dbConnect() (*mgo.Session, error) {
	return mgo.Dial(db.DbServer)
}

func (db *mongoDbDicProvider) getDicsCollection(session *mgo.Session) *mgo.Collection {
	return session.DB(db.DbName).C("dics")
}

func (db *mongoDbDicProvider) getMetaInfoCollection(session *mgo.Session) *mgo.Collection {
	return session.DB(db.DbName).C("metaInfo")
}

// Get list of registered dictionaries
func (db *mongoDbDicProvider) GetDictionaryList() ([]types.MetaInfoItem, error) {
	session, errConn := db.dbConnect()
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var metaInfo []types.MetaInfoItem
	var metaInfoCollection = db.getMetaInfoCollection(session)
	err := metaInfoCollection.Find(bson.M{}).All(&metaInfo)
	return metaInfo, err
}

//Get dictionary description
func (db *mongoDbDicProvider) GetDictionaryDesc(dicCode string) (*types.MetaInfoItem, error) {
	session, errConn := db.dbConnect()
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var metaInfo types.MetaInfoItem
	var metaInfoCollection = db.getMetaInfoCollection(session)
	err := metaInfoCollection.Find(nil).One(&metaInfo)
	return &metaInfo, err
}

/*
Get dictionary items
*/
func (db *mongoDbDicProvider) GetDictionaryItems(dicCode string) ([]types.DicItem, error) {
	session, errConn := db.dbConnect()
	if errConn != nil {
		log.Println("=== Db:GetDictionaryItems: failed to connect ")

		return nil, errConn
	}
	defer session.Close()

	var dicItems []types.DicItem = []types.DicItem{}
	var dicItemsCollection = db.getDicsCollection(session)
	err := dicItemsCollection.Find(bson.M{"dcode": dicCode}).All(&dicItems)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	return dicItems, err
}

/*
Get dictionary item by code
*/
func (db *mongoDbDicProvider) GetDictionaryItem(dicCode string, code string) (*types.DicItem, error) {
	session, errConn := db.dbConnect()
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var dicItem types.DicItem
	var dicItemsCollection = db.getDicsCollection(session)
	err := dicItemsCollection.Find(bson.M{"dcode": dicCode, "code": code}).One(&dicItem)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	return &dicItem, err
}
