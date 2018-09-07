package libs

import (
	"log"
	"gopkg.in/mgo.v2"
	"errors"
	"sync"
)

/*
 * Limitation: Only support one session for now
 * Todo:
 *	- handle errors
 *  - add mongodb config in env
 */

var logprefix = "[MONGODB]"

var minstance *mongo
var once sync.Once

// Singleton instance :D
func GetMongo() *mongo{
	once.Do(func() {
		minstance = &mongo{}
	})

	return minstance
}

type mongo struct {
	db *mgo.Database
	dbname string
}

// Establish a connection to MongoDb
func (mongo *mongo) Connect(dbname string) error {
	mongo.dbname = dbname
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		log.Println("Error", err)
		return errors.New(logprefix + err.Error())
	}

	log.Println("Connected to MongDB")
	mongo.db = session.DB(mongo.dbname)
	return nil
}

func (mongo *mongo) GetDB() (*mgo.Database) {

	if (mongo.db == nil) {
		mongo.Connect(mongo.dbname)
	}

	return mongo.db
}

