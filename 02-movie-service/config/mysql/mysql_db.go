package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// createMysqlDBConnection function for creating database connection
func createMysqlDBConnection(dsn string) *gorm.DB {
	sqlDB, err := sql.Open("mysql", dsn)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		defer sqlDB.Close()
		return gormDB
	}

	return gormDB
}

// LoadMysqlDB function for creating database connection from Mysql Database
func LoadMysqlDB() *gorm.DB {
	return createMysqlDBConnection(fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("MYSQL_DB_USERNAME"), os.Getenv("MYSQL_DB_PASSWORD"), os.Getenv("MYSQL_DB_HOST"), os.Getenv("MYSQL_DB_NAME")))
}
