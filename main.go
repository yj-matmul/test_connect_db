package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// loading password of DB
	var password string

	file, err := os.Open("./db_info.txt")
	if err != nil {
		log.Fatal("Cannot open db info file")
	}
	defer file.Close()
	fmt.Fscan(file, &password)

	// connect to a database
	_, err = sql.Open("pgx", fmt.Sprintf("host=localhost port=5001 dbname=test_connect user=postgres password=%s", password))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}

	// test my connection

	// get rows from table

	// insert a row

	// get rows from table again

	// update a row

	// get rows from table again

	// get one row by id

	// delete a row

	// get rosw again
}
