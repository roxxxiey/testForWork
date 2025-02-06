package main

import (
	"log"
	"os/exec"
)

func main() {
	log.Println("Start program ")
	cmd := exec.Command("./script.sh")

	log.Println("Command in cmd: ", cmd)

	err := cmd.Run()
	if err != nil {
		log.Print("Error executing script.sh: ", err)
		log.Fatal()
	}
}
