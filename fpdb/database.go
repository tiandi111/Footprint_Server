package mysql

import (
	"fmt"
//	"time"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

// Configuration
type Config struct {
	DNS		string
	Active		int
	Idle		int
	//IdleTimeout	time.Time
}

func NewMysql(c *Config) (db *sql.DB) {
	db, err := sql.Open("mysql", c.DNS)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to mysql!")
	}
	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	//db.SetConnMaxLifetime(time.Hour)
	return db
}

func CreateNewDatabase(db *sql.DB, name string) {
	cmd := "CREATE DATABASE IF NOT EXISTS "+name+";"
	_, err := db.Exec(cmd)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Created database "+name+" if not exists!")
	}
}

func CreateNewTable(db *sql.DB) {
	cmd := "CREATE TABLE IF NOT EXISTS User ( id INT AUTO_INCREMENT,name VARCHAR(255) NOT NULL, passwd VARCHAR(255) NOT NULL, email VARCHAR(255), PRIMARY KEY (id));"
	_, err := db.Exec(cmd)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Created main table User if not exists!")
	}
}











