package db

import (
	"database/sql"
	"fmt"
	"golang-rest-api/config"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	//connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_NAME + ")/" + conf.DB_PORT
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ")/" + conf.DB_NAME
	//"user:password@/dbname"
	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("Koneksi nya error")
	}

	err := db.Ping()

	if err != nil {
		//panic("Koneksinya kelamaan")
		fmt.Println("Koneksinyakelamaan ", err)
		os.Exit(1)
	}

}

func CreateCon() *sql.DB {
	return db
}
