package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Структура для хранения конфигурации БД
type dbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// Конфигурация БД
var dbCfg = dbConfig{
	User:     os.Getenv("DB_USER"),
	Password: os.Getenv("DB_PASSWORD"),
	Host:     os.Getenv("DB_HOST"),
	Port:     os.Getenv("DB_PORT"),
	Name:     os.Getenv("DB_NAME"),
}

// Вывод конфигурации БД
/*
func printDBConfig() {
	fmt.Println("DB_USER:", dbCfg.User)
	fmt.Println("DB_PASSWORD:", dbCfg.Password)
	fmt.Println("DB_HOST:", dbCfg.Host)
	fmt.Println("DB_PORT:", dbCfg.Port)
	fmt.Println("DB_NAME:", dbCfg.Name)
}
*/

var db *sql.DB

// Инициализация подключения к БД
func initDB() {
	time.Sleep(10 * time.Second)

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbCfg.User, dbCfg.Password,
		dbCfg.Host, dbCfg.Port,
		dbCfg.Name,
	)

	//printDBConfig()

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
		return
	}
	if err := db.Ping(); err != nil {
		log.Fatal("БД недоступна:", err)
		return
	}
	fmt.Println("Подключено к MySQL!")
}
