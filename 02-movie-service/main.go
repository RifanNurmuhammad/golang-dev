package main

import (
	"fmt"
	"os"
	"sync"

	config "github.com/joho/godotenv"
)

func main() {

	err := config.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	service := MakeHandler()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		service.HTTPServerMain()
	}()

	wg.Wait()
}
