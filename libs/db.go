package libs

import(
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var Dbconn *gorm.DB

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
		//log.Fatal("Error connecting to DB: %s", dbUri)
	}

	Dbconn = conn
	// automigrate
	//db.Debug().Automigrate()
}
