package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const(
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_users_schema = "mysql_users_schema"
)

var(
	Client *sql.DB

	username = 	os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	schema = os.Getenv(mysql_users_schema)
)

func init(){
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, host, schema)

	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil{
		panic(err)
	}
	if err = Client.Ping(); err != nil{
		panic(err)
	}

	log.Println("success db")
}