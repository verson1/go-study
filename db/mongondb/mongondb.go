package mongondb

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type mogConn struct {
	session *mgo.Session
}

func NewConnection() (MgoApi, error) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017"},
		Username: "admin",
		Password: "123456",
		Timeout:  time.Second * 30,
	})
	if err != nil {
		panic(err)
	}

	return &mogConn{session: session}, nil
}

func (m *mogConn) GetSession() *mgo.Session {
	return m.session
}

func (m *mogConn) GetDB(dbName string) *mgo.Database {
	return m.session.DB(dbName)
}

func (m *mogConn) Close() {
	m.Close()
}
