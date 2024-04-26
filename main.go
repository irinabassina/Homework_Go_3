package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Url struct {
	Name string
	Date time.Time
	Tags string
	Link string
}

func main() {
	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для добавления url нажмите клавишу 'a'")
	fmt.Println("Для удаления url нажмите клавишу 'r'")
	fmt.Println("Для просмотра всего списка url нажмите клавишу 'l'")
	fmt.Println("Для выхода из приложения нажмите Esc")
	urlCollection := make([]Url, 0)

OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {

		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильные аргументы в формате <url описание теги>")
				continue OuterLoop
			}

			// Напишите свой код здесь
			url := Url{args[1], time.Now(), args[2], args[0]}
			urlCollection = append(urlCollection, url)
			fmt.Println("Запись добавлена. Введите следующую команду")
		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>

			if len(urlCollection) > 0 {

				// Напишите свой код здесь
				for i, url := range urlCollection {
					fmt.Printf("Запись из списка URL адресов номер %d:\n", i)
					fmt.Printf("Имя: %s\n", url.Name)
					fmt.Printf("URL: %s\n", url.Link)
					fmt.Printf("Теги: %s\n", url.Tags)
					fmt.Printf("Дата: %s\n\n", url.Date.Format(time.DateTime))
				}
			} else {
				fmt.Println("Текущий список URL пуст")
			}
		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("Введите имя ссылки, которое нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			_ = text

			// Напишите свой код здесь
			idxToDelete := -1
			for i, url := range urlCollection {
				if url.Name == strings.TrimSpace(text) {
					idxToDelete = i
				}
			}
			if idxToDelete != -1 {
				urlCollection = append(urlCollection[:idxToDelete], urlCollection[idxToDelete+1:]...)
				fmt.Println("Ссылка удалена")
			} else {
				fmt.Println("Ссылка не найдена")
			}

		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
