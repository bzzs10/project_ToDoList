package todo

import "time"

type Event struct {
	InputText string    // Текст ошибки
	Time      time.Time // время ошибки
	ErrorText string    // описание ошибки
}

// файл со структурой пользовательского ввода(события), а так же важное уточнение, что сохраняться все события будут в формте слайса.
