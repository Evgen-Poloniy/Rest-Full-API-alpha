package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type dbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var dbCfg = dbConfig{
	User:     os.Getenv("DB_USER"),
	Password: os.Getenv("DB_PASSWORD"),
	Host:     os.Getenv("DB_HOST"),
	Port:     os.Getenv("DB_PORT"),
	Name:     os.Getenv("DB_NAME"),
}

var db *sql.DB

func initDB() {
	time.Sleep(10 * time.Second)

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbCfg.User, dbCfg.Password,
		dbCfg.Host, dbCfg.Port,
		dbCfg.Name,
	)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
		return
	}
	if err := db.Ping(); err != nil {
		log.Fatal("БД недоступна:", err)
		return
	}
	log.Printf("Подключено к MySQL!")
}
