package db

	import (
		"database/sql"
		"log"
		"os"
		"path/filepath"
	
		_ "github.com/mattn/go-sqlite3"
	)


	func ConnectDB() {
		appPath, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	
		// Определение переменной db вне блока if
		var db *sql.DB
	
		_, err = os.Stat(dbFile)
		if err != nil {
			// Присвоение значения переменной db внутри блока if
			db, err = sql.Open("sqlite3", dbFile)
			if err != nil {
				log.Fatalf("Ошибка при открытии базы данных: %v", err)
			}
			defer db.Close()
	
			// Создаем таблицу scheduler
			_, err = db.Exec(`CREATE TABLE scheduler (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				date DATE,
				title TEXT,
				comment TEXT,
				repeat_rule TEXT
			)`)
			if err != nil {
				log.Fatalf("Ошибка при создании таблицы: %v", err)
			}
	
			log.Println("База данных и таблица успешно созданы")
		} else {
			log.Println("Файл базы данных уже существует")
		}
	}
	