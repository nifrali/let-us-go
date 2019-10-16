package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SocietyOne/sqltest"
)

func main() {
	version, err := sqltest.TestDbConnection(
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("DATABASE"),
		os.Getenv("INSTANCE"),
		os.Getenv("PORT"),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)
}
