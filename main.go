package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// this section will ask for the contents of the empty cell
	// and the path to the input csv
	// the nan can be empty
	// the csv cannot.,,,;./
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	fmt.Println(arg1)
	fmt.Println(arg2)

	rows := readSample()
	appendSum(rows)
	writeChanges(rows)
}

func readSample() [][]string {
	// f, err := os.Open("sample.csv")
	f, err := os.Open("sample3.csv")
	// f, err := os.Open("interpolated_test_data.csv")
	//f, err := os.Open("input_test_data.csv")

	if err != nil {
		log.Fatal(err)
	}
	rows, err := csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func appendSum(rows [][]string) {
	// change the name and do not need sum but might use it to get the average
	// walk through the rows
	// when you come to a stop-word then get the two horizontal adjacent cells
	// and the two vertical adjacent cells
	// get the average and inset it into the current cell
	rows[0] = append(rows[0], "SUM")
	for i := 1; i < len(rows); i++ {
		fmt.Println(rows[i][0])
		fmt.Println("---")
		fmt.Println(len(rows))
		rows[i] = append(rows[i], sum(rows[i]))
	}
}

func sum(row []string) string {
	sum := 0
	for _, s := range row {
		x, err := strconv.Atoi(s)
		if err != nil {
			return "NA"
		}
		sum += x
	}
	return strconv.Itoa(sum)
}

func writeChanges(rows [][]string) {
	f, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = csv.NewWriter(f).WriteAll(rows)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
