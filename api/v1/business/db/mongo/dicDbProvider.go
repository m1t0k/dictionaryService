package db

import (
	AppConfig "../infrastructure/config"
	DicTypes "../types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
EventDBProvider type
*/
type dicDbProvider struct{}

func dbConnect() (*mgo.Session, error) {
	return mgo.Dial(AppConfig.AppConfiguration.DbServer)
}

func getDicsCollection(session *mgo.Session) *mgo.Collection {
	return session.DB(AppConfig.AppConfiguration.DbName).C("dics")
}

func getMetaInfoCollection(session *mgo.Session) *mgo.Collection {
	return session.DB(AppConfig.AppConfiguration.DbName).C("metaInfo")
}

// Get list of registered dictionaries
func (db *dicDbProvider) GetDictionaryList() ([]DicTypes.MetaInfoItem, error) {
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
func (db *dicDbProvider) GetDictionaryDesc(dicCode string) ([]DicTypes.MetaInfoItem, error) {
	session, errConn := dbConnect()
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var metaInfo []DicTypes.MetaInfoItem
	var metaInfoCollection = getMetaInfoCollection(session)
	err := metaInfoCollection.Find(bson.M{Code: dicCode}).All(&metaInfo)
	return metaInfo, err
}

/*
Get dictionary items
*/
func (db *dicDbProvider) GetDictionaryItems(dicCode string) ([]DicTypes.DicItem, error) {
	session, errConn := dbConnect()
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var dicItems []DicTypes.DicItem
	var dicItemsCollection = getDicsCollection(session)
	err := dicItemsCollection.Find(bson.M{DicCode: dicCode}).All(&dicItems)
	return dicItems, err
}

/*
Get dictionary item by code
*/
func (db *dicDbProvider) GetDictionaryItem(dicCode string, code string) ([]DicTypes.DicItem, error) {
	session, errConn := dbConnect()
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var dicItem DicTypes.DicItem
	var dicItemsCollection = getDicsCollection(session)
	err := dicItemsCollection.Find(bson.M{DicCode: dicCode, Code: code}).One(&dicItem)
	return dicItem, err
}
