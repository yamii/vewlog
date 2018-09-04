package  models

import(
	"crypto/bcrypt"
	"github.com/jinzhu/gorm"

	"vewlog/libs"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Username string `json:"username`
	Password string `json:"password"`
}


func (user *User) Create() (map[string] interface{}) {
	hashedPassword, _ := brcypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.password = string(hashedPassword)

	libs.GetDb().Create(account)
}
