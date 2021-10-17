package mongondb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

var conn MgoConnect

type Person struct {
	Name   string `bson:"name"`
	Age    int    `bson:"age"`
	Gender string `bson:"gender"`
}

func init() {
	var err error
	conn, err = NewConnection(
		WithUrl("mongodb://127.0.0.1:27017"),
		WithMechanism("SCRAM-SHA-1"),
		WithUsernameAndPassword("root", "123456"),
		WithDatabase("admin"),
	)
	if err != nil {
		panic(err)
	}
}

func TestInsert(t *testing.T) {
	db := conn.GetDB("test")
	// 创建collection
	//err := db.CreateCollection(context.Background(),"users")
	//if err != nil {
	//	panic(err)
	//}
	result, err := db.Collection("users").InsertOne(context.Background(), bson.M{"name": "李四", "gender": "男", "age": 18})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestInsertMany(t *testing.T) {
	db := conn.GetDB("test")
	pes := []interface{}{Person{
		Name:   "王五",
		Gender: "男",
		Age:    17,
	}, Person{
		Name:   "早起",
		Gender: "女",
		Age:    25,
	},
	}

	res, err := db.Collection("users").InsertMany(context.Background(), pes)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestFind(t *testing.T) {
	db := conn.GetDB("test")
	//res,err := db.Collection("users").Find(context.Background(),bson.M{})
	//if err != nil {
	//	panic(err)
	//}
	// 查询age > 14, 查询2条数据，按降序排序
	res, err := db.Collection("users").Find(context.Background(), bson.M{"age": bson.M{"$gte": 14}},
		options.Find().SetLimit(2), options.Find().SetSort(bson.M{"age": -1}))
	if err != nil {
		panic(err)
	}
	var mymap []Person
	err = res.All(context.Background(), &mymap)
	if err != nil {
		panic(err)
	}
	fmt.Println(mymap)
}

func TestUpdate(t *testing.T) {
	db := conn.GetDB("test")

	res, err := db.Collection("users").UpdateOne(context.Background(), bson.M{"name": "张三"}, bson.M{"$set": bson.M{"age": 43}})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
