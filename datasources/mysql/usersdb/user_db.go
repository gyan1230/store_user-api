package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	//added mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	//Client :
	Client *sql.DB
)

const (
	username = "root"
	password = "root"
	hostname = "localhost:3306"
	schema   = "users_db"
)

//Connect :
func Connect() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, hostname, schema)
	Client, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	if err := Client.Ping(); err != nil {
		log.Panicln(err)
		return nil
	}
	fmt.Println("DB connected sucessfully")
	return Client
}

// func init() {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, hostname, schema)
// 	Client, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Panicln(err)

// 	}
// 	if err := Client.Ping(); err != nil {
// 		log.Panicln(err)

// 	}
// 	fmt.Println("DB connected sucessfully")

// }
