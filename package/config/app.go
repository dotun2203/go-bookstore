package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// func Connect() {
// 	dsn := "root:adedotun2203@localhost:3306/bookstore"
// 	data, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer data.Close()

// 	err = data.Ping()
// }

// func GetDB() *gorm.DB {
// 	return db
// }

func Connect() {
	db, err := sql.Open("mysql", "root:adedotun2203@localhost:3306/bookstore")
	if err != nil {
		fmt.Println("error validating sql.open arguements")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.ping")
		panic(err.Error())
	}
}
