package core

import (
	"context"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	session *mgo.Session
	mClient *mongo.Client
	dbname  string
)

func getDbName(uri string) string {
	if uri == "" {
		return ""
	}

	sl := strings.LastIndex(uri, "/")
	if sl < 0 {
		return ""
	}

	n := strings.TrimPrefix(uri[sl:], "/")
	i := strings.Index(n, "?")
	if i > 0 {
		n = n[:i]
	}
	return n
}

func InitDb() {
	url := beego.AppConfig.String("mongodb::url")

	dbname = getDbName(url)
	if dbname == "" {
		panic("database name is invalid")
	}

	initmgoDB(url)
}

func initmgoDB(url string) {
	var err error
	session, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	err = session.Ping()
	if err != nil {
		panic(err)
	}
	//session = sess
	session.SetMode(mgo.Monotonic, true)
	//session.SetMode(mgo.Eventual, true)
	session.SetPoolLimit(2000)
	session.SetSocketTimeout(19 * time.Second)

	//支持事务
	mClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
}

// Conn return mongodb session.
func GetmgoConn() *mgo.Session {
	return session.Copy()
}

//获取数据库的collection
func GetmgoColl(name string) *mgo.Collection {
	return session.DB(dbname).C(name)
}

func GetmgoDB() *mgo.Database {
	return session.DB(dbname)
}

func GetmgoCollT(name string) *mongo.Collection {
	return mClient.Database(dbname).Collection(name)
}

func GetmgoDBT() *mongo.Database {
	return mClient.Database(dbname)
}

// func GetDBClient() *mongo.Client {
// 	return mClient.Database(dbname).Client()
// }
