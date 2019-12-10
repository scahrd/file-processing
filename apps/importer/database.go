package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func testDbConnection() {
	fmt.Println("Checking database connection ...")
	db := connect()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected successfully!")
}

func connect() *sql.DB {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	return db
}

func insert(table string, data []string) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (cpf, privado, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_frequente, loja_ultima_compra) VALUES (%s)",
		table, strings.Join(data, ", "))

	db := connect()
	defer db.Close()

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Print(".")
	return nil
}
