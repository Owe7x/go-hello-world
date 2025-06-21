package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var message string
	message = hello("Максим", 21, 34)

	fmt.Println(message)
}

func hello(name string, age int, age2 int) string {
	return fmt.Sprintf("Привет. %s! Тебе %d лет %d", name, age, age2)
}
