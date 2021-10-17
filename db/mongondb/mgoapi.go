package mongondb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MgoConnect interface {
	GetClient() *mongo.Client
	GetDB(dbName string) *mongo.Database
	Close()
}
