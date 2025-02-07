package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Start program")

	// Получаем полный путь к скрипту
	scriptPath := "./pingNAP"

	// Создаём команду с явным указанием shell
	cmd := exec.Command("sudo", scriptPath)

	// Перенаправляем stdout и stderr на консоль
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Запускаем скрипт
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing script.sh:", err)
		os.Exit(1)
	}

	fmt.Println("End program")
}
