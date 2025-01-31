package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	//"os/user"

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
func printDBConfig() {
	fmt.Println("DB_USER:", dbCfg.User)
	fmt.Println("DB_PASSWORD:", dbCfg.Password)
	fmt.Println("DB_HOST:", dbCfg.Host)
	fmt.Println("DB_PORT:", dbCfg.Port)
	fmt.Println("DB_NAME:", dbCfg.Name)
}

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

// Структура для хранения записи
type Record struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
	Time    string  `json:"last_time"`
}

func responseRecord(w http.ResponseWriter, record Record) {
	w.Header().Set("Content-Type", "application/json")

	responseJSON, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка формирования JSON", http.StatusInternalServerError)
		return
	}

	w.Write(append(responseJSON, '\n'))
}

// Структура для вывода ошибки в формате JSON
type ErrorResponse struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

// Вывод кода ошибки в формате JSON
func responseError(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Error:   statusCode,
		Message: errMsg,
	}

	responseJSON, _ := json.MarshalIndent(response, "", "  ")

	w.Write(append(responseJSON, '\n'))
}

// Получение записи по ID
func getRecordByID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		responseError(w, 400, "Параметр id обязателен")
		return
	}

	var record Record
	err := db.QueryRow("SELECT id, balance, last_time FROM users WHERE id = ?", userID).Scan(&record.ID, &record.Balance, &record.Time)

	if err != nil {
		responseError(w, 404, "Пользователь не найден")
		return
	}

	responseRecord(w, record)
}

func getCountOfUsers(w http.ResponseWriter, r *http.Request) {
	var count int
	err := db.QueryRow("SELECT COUNT(DISTINCT id) FROM users").Scan(&count)
	if err != nil {
		log.Println("Ошибка запроса к БД:", err)
		http.Error(w, "Ошибка запроса к базе данных", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Количество пользователей: %d\n", count)
}

func main() {
	initDB()
	http.HandleFunc("/getRecordByID", getRecordByID)
	http.HandleFunc("/getCountOfUsers", getCountOfUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
