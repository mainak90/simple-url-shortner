package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"os"
)

func InitDB() (*sql.DB, error) {
	pgUrl, err := pq.ParseURL(os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db, err := sql.Open("postgres", pgUrl)
	if err != nil {
		return nil, err
	} else {
		stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS WEB_URL(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		_, err = stmt.Exec()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return db, nil
	}
}
