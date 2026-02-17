package main

import "time"

type Task struct {
	Title       string    // заголовок (1 слово)
	Description string    // основной текст задачи
	CreatedAt   time.Time // время создания
	Done        bool      // выполнена или нет
	DoneAt      time.Time // время выполнения, если выполнена
}

// файл со структурой мапы для заданий
