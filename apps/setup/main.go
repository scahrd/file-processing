package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("BEGINNING SETUP!")
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Error loading .env file")
	}
}

func readSQL(queryName string) string {
	file, err := os.Open("setup/sql/" + queryName + ".sql")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result, _ := ioutil.ReadAll(file)
	return (string(result))
}

func connect() *sql.DB {
	fmt.Println("Attempting to connect to database ...")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Println("Failed to connect!")
		panic(err)
	}
	fmt.Println("Connection was successfull!")
	return db
}

func main() {
	db := connect()
	defer db.Close()

	_, errModels := db.Exec(readSQL("models"))
	if errModels != nil {
		panic(errModels)
	}
	fmt.Println("Table created!")

	_, errTriggers := db.Exec(readSQL("triggers"))
	if errTriggers != nil {
		panic(errTriggers)
	}
	fmt.Println("Triggers created!")
	fmt.Println("SETUP DONE!")

}
