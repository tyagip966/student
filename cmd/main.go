package main

import (
	"log"
	"student/utils"
)

func main() {
	container := utils.Container{Injected: false}
	err := container.TriggerDI()
	if err != nil {
		log.Fatal("Error starting Server")
	}
}
