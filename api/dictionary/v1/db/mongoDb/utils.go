package mongoDb

import mgo "gopkg.in/mgo.v2"

func Connect(server string) (*mgo.Session, error) {
	return mgo.Dial(server)
}

func GetCollection(session *mgo.Session, dbName string, colName string) *mgo.Collection {
	return session.DB(dbName).C(colName)
}
