package mongondb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type mgoConn struct {
	client *mongo.Client
}

func (m mgoConn) GetClient() *mongo.Client {
	return m.client
}

func (m mgoConn) GetDB(dbName string) *mongo.Database {
	return m.client.Database(dbName)
}

func (m mgoConn) Close() {
	err := m.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}

func NewConnection(opt ...Option) (MgoConnect, error) {
	opts := &Options{}
	for _, o := range opt {
		o(opts)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(opts.URI).
			SetAuth(options.Credential{
				AuthMechanism: opts.Mechanism,
				AuthSource:    opts.Database,
				Username:      opts.Username,
				Password:      opts.Password,
			}),
	)
	if err != nil {
		return nil, err
	}
	return &mgoConn{client: client}, nil
}
