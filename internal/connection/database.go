package connection

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	"log"
	"database/sql"
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
		log.Fatal("Error connecting to the database:", err.Error())
	}

	err = db.Ping()
	if err != nil {	
		log.Fatal("Error pinging the database:", err.Error())
	}

	return db
}