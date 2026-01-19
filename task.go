package main

import (
	"fmt"
	"log"
)

type Task struct {
	Text      string
	Completed bool
}

func addTask(task string) error {
	if task == "" {
		return fmt.Errorf("Задача не может быть пустой!")
	}
	return nil
}

func main() {
	var tasks []Task
	var choice int

	for {
		fmt.Print("\n--To-Do--")

		fmt.Println("\n\n--1.Добавить задачу--")

		fmt.Println("--2.Просмотреть задачи--")

		fmt.Println("--3.Пометить выполненой--")

		fmt.Println("--0.Выход--")

		fmt.Print("\n\nВыберите действие: ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			var text string

			fmt.Println("\nДобавьте задачу: ")
			fmt.Scanln(&text)

			err := addTask(text)
			if err != nil {
				log.Println("Ошибка!", err)
				continue
			}

			tasks = append(tasks, Task{
				Text:      text,
				Completed: false,
			})

			log.Println("\n\nЗадача добавлена")
			fmt.Println("Задача сохранена")

		case 2:

			if len(tasks) == 0 {
				fmt.Println("\nЗадача отсутствует")
				continue
			}

			for i, task := range tasks {
				status := "Не выполнено"
				if task.Completed {
					status = "Выполнено"
				}
				fmt.Println(i, task.Text, status)
			}

		case 3:

			if len(tasks) == 0 {
				fmt.Println("\nЗадачи нет")
				continue
			}

			var index int

			fmt.Println("\nВведите номер задачи: ")
			fmt.Scanln(&index)

			if index < 0 || index >= len(tasks) {
				log.Println("\nНеверный номер задачи")
				fmt.Println("\nТакой задачи нет")
				continue
			}

			tasks[index].Completed = true
			log.Println("\nЗадача отмечена выполненной")
			fmt.Println("Задача выполнена")

		case 0:
			log.Println("\nВыход...")
			fmt.Println("Покааа...")
			return

		default:
			log.Println("\nНе верная команда")
			fmt.Println("Выбрано неверное действие")
		}
	}
}
