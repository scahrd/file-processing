package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Error loading .env file")
	}
}
func main() {
	fmt.Println("Initializing ...")
	testDbConnection()
	files, err := getFiles("importer/files/")
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		fmt.Println("You don't have any file to process.")
		os.Exit(0)
	}

	for _, file := range files {
		fmt.Println("... Starting file processing ...")
		processFile(file)
	}

}
