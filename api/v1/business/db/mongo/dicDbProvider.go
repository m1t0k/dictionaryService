package db

import (
	DicTypes "../../types/"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
EventDBProvider type
*/
type MongoDbDicProvider struct{}

func dbConnect() (*mgo.Session, error) {
	return mgo.Dial("dbServer")
}

func getDicsCollection(session *mgo.Session) *mgo.Collection {
	return session.DB("dbName").C("dics")
}

func getMetaInfoCollection(session *mgo.Session) *mgo.Collection {
	return session.DB("dbName").C("metaInfo")
}

// Get list of registered dictionaries
func (db *MongoDbDicProvider) GetDictionaryList() ([]DicTypes.MetaInfoItem, error) {
	session, errConn := dbConnect()
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var metaInfo []DicTypes.MetaInfoItem
	var metaInfoCollection = getMetaInfoCollection(session)
	err := metaInfoCollection.Find(bson.M{}).All(&metaInfo)
	return metaInfo, err
}

//Get dictionary description
func (db *MongoDbDicProvider) GetDictionaryDesc(dicCode string) (DicTypes.MetaInfoItem, error) {
	session, errConn := dbConnect()
	if errConn != nil {
		return DicTypes.MetaInfoItem{}, errConn
	}
	defer session.Close()

	var metaInfo DicTypes.MetaInfoItem
	var metaInfoCollection = getMetaInfoCollection(session)
	err := metaInfoCollection.Find(nil).One(&metaInfo)
	return metaInfo, err
}

/*
Get dictionary items
*/
func (db *MongoDbDicProvider) GetDictionaryItems(dicCode string) ([]DicTypes.DicItem, error) {
	session, errConn := dbConnect()
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var dicItems []DicTypes.DicItem
	var dicItemsCollection = getDicsCollection(session)
	err := dicItemsCollection.Find(nil).All(&dicItems)
	return dicItems, err
}

/*
Get dictionary item by code
*/
func (db *MongoDbDicProvider) GetDictionaryItem(dicCode string, code string) (DicTypes.DicItem, error) {
	session, errConn := dbConnect()
	if errConn != nil {
		return DicTypes.DicItem{}, errConn
	}
	defer session.Close()

	var dicItem DicTypes.DicItem
	var dicItemsCollection = getDicsCollection(session)
	err := dicItemsCollection.Find(nil).One(&dicItem)
	return dicItem, err
}
