package main

import (
	"fmt"
	"go_final_project/internal/db"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

db.ConnectDB()

webDir := "./web"
// Устанавливаем обработчик для маршрута "/"
http.Handle("/", http.FileServer(http.Dir(webDir)))

// Запускаем веб-сервер на порту 7540
fmt.Println("Сервер запущен на порту 7540")
http.ListenAndServe(":7540", nil)
}


