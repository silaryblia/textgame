package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	initGame()
	fmt.Println("Добро пожаловать в текстовую игру! (введите 'выход' чтобы завершить)")
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !reader.Scan() {
			break
		}
		cmd := reader.Text()
		if cmd == "выход" {
			fmt.Println("Игра завершена.")
			break
		}
		fmt.Println(handleCommand(cmd))
	}
}

///
