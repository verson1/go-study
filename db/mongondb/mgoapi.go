package mongondb

import "gopkg.in/mgo.v2"

type MgoApi interface {
	GetSession() *mgo.Session
	GetDB(dbName string) *mgo.Database
	Close()
}
