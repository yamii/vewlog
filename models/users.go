package  models

import(
    //"log"
    "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"

    libs "vewlog/libs"
)

const (
    collection = "users"
)

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
    Name     string `json:name binding:"-"`
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func (user *User) userIndex() mgo.Index {
	return mgo.Index {
		Key        : []string{"username"},
		Unique     : true,
		DropDups   : true,
		Background : true,
		Sparse     : true,
	}
}

func userModelIndex() mgo.Index {
    return mgo.Index {
        Key        : []string{"username"},
        Unique     : true,
        DropDups   : true,
        Background : true,
        Sparse     : true,
    }
}

/*
 * @user User struct
*/
func (user *User) Create()error {
    table := libs.GetMongo().GetDB().C(collection)
    table.EnsureIndex(userModelIndex())

    err := table.Insert(&user)
    return err
}
