package main

import (
	"database/sql" // Стандартный инструмент Go для работы с базами данных
	"fmt"
	"log"

	_ "modernc.org/sqlite" // Подключаем драйвер SQLite (он зарегистрирует себя сам)
)

func main() {
	// 1. Подключаемся к файлу твоей базы данных school.db
	db, err := sql.Open("sqlite", "school.db")
	if err != nil {
		log.Fatal("Не удалось открыть файл базы данных:", err)
	}
	defer db.Close() // Закрываем соединение, когда программа закончит работу

	// 2. Включаем контроль связей (наш любимый PRAGMA), чтобы SQLite проверяла ключи
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("Не удалось включить PRAGMA:", err)
	}

	// 3. Пишем SQL-запрос для поиска учеников из 5 "А" класса
	query := `
		SELECT students.full_name, students.parent_phone 
		FROM students 
		JOIN classes ON students.class_id = classes.class_id 
		WHERE classes.grade_year = 5 AND classes.letter = 'А';
	`

	// 4. Отправляем запрос в базу данных
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса:", err)
	}
	defer rows.Close() // Очищаем память после чтения строк

	fmt.Println("=== Список учеников 5 'А' класса ===")
	fmt.Println("----------------------------------------")

	// Переменная-флаг, чтобы проверить, нашли ли мы кого-нибудь
	found := false

	// 5. Читаем результаты построчно в цикле
	for rows.Next() {
		found = true
		var fullName string
		var parentPhone string

		// Scan раскладывает ФИО и телефон из таблицы по нашим переменным в Go
		err := rows.Scan(&fullName, &parentPhone)
		if err != nil {
			log.Fatal("Ошибка чтения строки:", err)
		}

		// Выводим ученика на экран
		fmt.Printf("Ученик: %s | Телефон родителя: %s\n", fullName, parentPhone)
	}

	if !found {
		fmt.Println("Ученики в 5 'А' классе пока не найдены. Добавь их в базу!")
	}

	// Проверяем, не оборвалось ли чтение из-за ошибки в процессе
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}