package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	in, err := os.ReadFile("./data.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(strings.NewReader(string(in)))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}
