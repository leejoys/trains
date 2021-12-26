package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func howLong(t1, t2 string) (int, error) {

}

func cheapest(records [][]string) ([]string, error) {
	g := map[Name]Neighbors{}

	for _, record := range records {
		price, err := strconv.Atoi(strings.Trim(record[3], "."))
		if err != nil {
			return nil, err
		}
		if _, ok := g[record[1]]; !ok {
			g[Name(record[1])] = Neighbors{}
		}
		err = g.AddVertex(record[1])
		if err != nil {
			return nil, err
		}
		err = g.AddVertex(record[2])
		if err != nil {
			return nil, err
		}
		err = g.AddEdge(record[1], record[2])
		if err != nil {
			return nil, err
		}
	}

}

func fastest(records [][]string) {

}

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
	cheapest(records)
	fastest(records)
}
