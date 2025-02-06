package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("./script.sh")

	err := cmd.Run()
	if err != nil {
		log.Print("Error executing script.sh: ", err)
		log.Fatal()
	}
}
