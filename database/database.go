package database

import (
	"database/sql"
	"fmt"
	"golang-rest-api/config"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	//connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_NAME + ")/" + conf.DB_PORT
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ")/" + conf.DB_NAME
	//"user:password@/dbname"
	database, err = sql.Open("mysql", connectionString)
	database.SetMaxOpenConns(1000)

	if err != nil {
		panic("Opps, get something wrong with connections.")
	}

	err := database.Ping()

	if err != nil {
		fmt.Println("Connection too long to handle request.", err)
		os.Exit(1)
	}


}

func CreateCon() *sql.DB{
	return database
}

//digunakan untuk closing setiap koneksi yang pernah terjadi
//agar tidak terjadi stuck akses
func CloseDatabase() {
  	database.Close()
}
