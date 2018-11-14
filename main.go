package main

import (
	"database/sql"
	"os"
	"log"
	"io/ioutil"
	"strings"

	//Extends database/sql to support postgres
	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("POSTGRES_URL")
	if len(connStr) == 0 {
		log.Fatal("POSTGRES_URL not set, exiting")
	}

	file, err := ioutil.ReadFile("commands.sql")

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	strFile := string(file)
	strSplit := strings.Split(strFile, ";")

	for _,str := range strSplit{
		if str != "\n" {
			_, err = db.Query(str)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	db.Close()

	log.Println("Successfully ran SQL query")
}
