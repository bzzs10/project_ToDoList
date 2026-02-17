package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("введите команду: ")
		ok := scanner.Scan()
		if !ok {
			fmt.Println("Ошибка ввода!")
			return
		}

		text := scanner.Text()

		Fields := strings.Fields(text)
		if len(Fields) == 0 {
			continue
		}

		cmd := Fields[0] // первое слово которое читает программа

		if cmd == "выйти" {
			fmt.Println("вы завершили программу!")
			return
		}

		if cmd == "help" {
			fmt.Println("список команд...")
			continue
		}

		if cmd == "добавить" {
			if len(Fields) != 5 {
				fmt.Println("неверный формат команды")
				continue
			}
		}
	}
}

// пересмотреть пользовательский ввод, и сделать в store функции, и через main обрашаться к функция по типу удалить, добавить и т.д.
// после придумать что то, что будет сохронять события в store, скорее функция, и все сделать через условные вветвелния чтобы получать
// либо сохраненную задачу/событие либо ошибку(к примеру: неправельный формат команды или такой задачи нету)
