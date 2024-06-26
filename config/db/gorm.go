package db

import (
	"fmt"
	"os"
	
	"github.com/jinzhu/gorm"
)

type SqlHandler struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
}

func NewDB() *SqlHandler {
	readConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_READ_USER"),
		os.Getenv("MYSQL_READ_PASSWORD"),
		os.Getenv("MYSQL_READ_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
    readDB, err := gorm.Open("mysql", readConn)
    if err != nil {
        panic(err.Error())
    }
	readDB.SingularTable(true)

	writeConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_WRITE_USER"),
		os.Getenv("MYSQL_WRITE_PASSWORD"),
		os.Getenv("MYSQL_WRITE_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
    writeDB, err := gorm.Open("mysql", writeConn)
    if err != nil {
        panic(err.Error())
    }
	writeDB.SingularTable(true)

    sqlHandler := SqlHandler{
        ReadConn:  readDB,
        WriteConn: writeDB,
    }

    return &sqlHandler
}
