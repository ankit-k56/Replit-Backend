package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
var Db *gorm.DB

func ConnecttoDb() {
	var err error
	
	dsn := os.Getenv("DATABASE_URL")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if(err != nil){
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to Database")

}