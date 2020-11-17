package database

import (
	"log"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

var Sess db.Session

var settings = postgresql.ConnectionURL{
	Database: `corporative_system`,
	Host:     `localhost`,
	User:     `postgres`,
	Password: `docker`,
}

func Initalize() {

  sess, err := postgresql.Open(settings)
	
	if err != nil {
  	log.Fatal(err)
	}
	
	Sess = sess

}


func Close() {
	Sess.Close()
}