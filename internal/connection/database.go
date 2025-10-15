package connection

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	"log"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
)

func GetDatabase(conf config.Database) *sql.DB{ 
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Tz,
	)

	db, err := sql.Open("postgres", dsn)	

	if err != nil {
		utility.CreateLog("warn", fmt.Sprintf("Failed connect to database: %v", err.Error()), "application")
		log.Fatal("Failed connect to database:", err.Error())
	}

	err = db.Ping()
	if err != nil {	
		utility.CreateLog("warn", fmt.Sprintf("Error pinging the database:: %v", err.Error()), "application")
		log.Fatal("Error pinging the database:", err.Error())
	}

	return db
}