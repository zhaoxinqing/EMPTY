package lib

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// db连接var db *gorm.DB// Setup 初始化连接
func Setup() { // db = newConnection()

	dialector := mysql.New(mysql.Config{
		DSN:                       "dbURI", // data source name
		DefaultStringSize:         256,     // default size for string fields
		DisableDatetimePrecision:  true,    // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,    // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,    // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,   // auto configure based on currently MySQL version
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Print(err.Error())
	}

	sqlDB, err := conn.DB()
	if err != nil {
		fmt.Println("connect db server failed.")
	}
	sqlDB.SetMaxIdleConns(10)                   // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(100)                  // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Second * 600) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db = conn
}

// GetDB 开放给外部获得db连接
func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("connect db server failed.")
		Setup()
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		Setup()
	}
	return db
}

// https://bbs.huaweicloud.com/blogs/200862
