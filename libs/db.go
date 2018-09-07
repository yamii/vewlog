package libs

import(
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

func InitDB() {
	user     := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name     := os.Getenv("DB_NAME")
	host     := os.Getenv("DB_HOST")

	dbUri := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		host, user, name, password,
	)

	log.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Println("Error connecting to DB: %s", dbUri)
	}

	conn = conn
}

func GetDB() *gorm.DB {
	// get connection if no connection aquired
	if conn == nil {
		InitDB()
	}

	return conn
}
