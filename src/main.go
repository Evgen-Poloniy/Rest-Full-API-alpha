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

func printDBConfig() {
	fmt.Println("DB_USER:", dbCfg.User)
	fmt.Println("DB_PASSWORD:", dbCfg.Password)
	fmt.Println("DB_HOST:", dbCfg.Host)
	fmt.Println("DB_PORT:", dbCfg.Port)
	fmt.Println("DB_NAME:", dbCfg.Name)
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

	printDBConfig()

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

type Record struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
	Time    string  `json:"last_time"`
}

// Создание записи
func createRecord(w http.ResponseWriter, r *http.Request) {
	var record Record

	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	if record.ID != 0 {
		_, err = db.Exec("INSERT INTO users (id, balance) VALUES (?, ?)", record.ID, record.Balance)
	} else {
		_, err = db.Exec("INSERT INTO users (balance) VALUES (?)", record.Balance)
	}

	if err != nil {
		http.Error(w, "Ошибка записи в БД", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

// Получение записи по ID
func getRecordByID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, `{"status": "error", "message": "ID пользователя не указан"}`, http.StatusBadRequest)
		return
	}

	var record Record
	err := db.QueryRow("SELECT id, balance, time FROM users WHERE id = ?", userID).Scan(&record.ID, &record.Balance, &record.Time)
	if err != nil {
		http.Error(w, `{"status": "error", "message": "Пользователь не найден"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func getCountOfUsers(w http.ResponseWriter, r *http.Request) {
	var count int
	// Выполняем запрос к БД
	err := db.QueryRow("SELECT COUNT(DISTINCT id) FROM users").Scan(&count)
	if err != nil {
		// Логируем ошибку и отправляем ответ с ошибкой в тело ответа
		log.Println("Ошибка запроса к БД:", err)
		http.Error(w, "Ошибка запроса к базе данных", http.StatusInternalServerError)
		return
	}
	// Отправляем результат пользователю
	fmt.Fprintf(w, "Количество пользователей: %d\n", count)
}

func main() {
	initDB()
	http.HandleFunc("/createRecord", createRecord)
	http.HandleFunc("/getRecordByID", getRecordByID)
	http.HandleFunc("/getCountOfUsers", getCountOfUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
