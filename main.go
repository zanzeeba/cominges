package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// this section will ask for the contents of the empty cell
	// and the path to the input csv
	// the nan can be empty
	// the csv cannot

	var fp = flag.String("fp", "", "Pass in the filename and path (Required)")
	var altnan = flag.String("altnan", "nan", "The default indicator of a missing number is nan but you can add your own string here (Optional)")
	var help = flag.String("h", "", "Extended help")
	flag.Parse()

	if *fp == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("flag command line:- %+v\n", *fp)
	fmt.Printf("flag command line:- %+v\n", *altnan)
	fmt.Printf("flag command line:- %+v\n", *help)

	rows := readSample()
	appendSum(rows)
	writeChanges(rows)
}

func readSample() [][]string {
	// f, err := os.Open("sample.csv")
	// f, err := os.Open("sample3.csv")
	// f, err := os.Open("interpolated_test_data.csv")
	f, err := os.Open("./input_data/input_test_data.csv")

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

	// get the size of the array
	// row
	mr := len(rows)
	fmt.Printf("max rows is:- %+v\n", mr)
	// length of row all will be the same so just get the first one
	lr := len(rows[0])
	fmt.Printf("length of row is:- %+v\n", lr)

	// if the row/col is 0/0 then can only get right side and below
	// if the row/col is 0/1-len-1 either side and below
	// if the row/col is 0/len then left side and below

	// if the row/col is 1 to depth -1/0 then above, below, right
	// if the row/col is 1 to depth -1/1-len-1 then above, below, left, right
	// if the row/col is 1 to depth -1/len then above, below, left

	// if the row/col is max depth/0 then above, right
	// if the row/col is max depth/1-len-1 then above, left, right
	// if the row/col is max depth/len then above, left

	rows[0] = append(rows[0], "SUM")
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j] == "nan" {
				fmt.Printf("hit a not a number %+v row is %+v and col is %+v\n", rows[i][j], i, j)
				newVal := CalcAverage(1.1, 2.1, 3.1, 4.1)
				fmt.Println(newVal)
				// mr = max rows = the number of rows(calculated after reading in the file)
				// rl = row length  (calculated after reading in the file)
				// cr = current row, pr = previous row, nr - next row
				// cc = current column, pc = previous column, nc = next column
				// if row == 0
				// if col is 0
				// avg of row[cr,1] + row[nr,0]
				// if 1 to rl-1
				// avg of row[cr,cc-1] + row[cr,cc+1] + row[nr,cc]
				// if col is rl
				// avg of row[cr, rl-1] + row[nr,rl]
				//row 1 to mr
				// if col is 0
				// avg of row[cr,cc+1] + row[pr,cc]  + row[nr,cc]
				// if col is 1 to rl-1
				// avg of row[cr,cc-1] + row[cr,cc+1] + row[pr,cc]  + row[nr,cc]
				// if col is rl
				// avg of row[cr/cc-1] + row[pr,cc]  + row[nr,cc]
				// if row == mr
				// if col is 0
				// avg of row[cr,1] + row[pr,0]
				// if cols 1 to rl-1
				// avg of row[cr,cc-1] + row[cr,cc+1] + row[pr,cc]
				// if col is rl
				// avg of row[cr, rl-1] + row[pr,rl]

				//fmt.Println(rows[i][j])
			}
		}
		fmt.Println("------------")
		//fmt.Println(rows[i][0])
		//fmt.Println("---")
		//fmt.Println(len(rows))
		rows[i] = append(rows[i], sum(rows[i]))
	}
}

func CalcAverage(nums ...float64) float64 {
	//fmt.Print(nums, " ")
	//fmt.Println(len(nums))
	var total float64
	for _, num := range nums {
		total += num
	}
	//fmt.Println(total)
	totavg := total / float64(len(nums))

	fmt.Printf("the numbers passed was %+v\n", len(nums))
	fmt.Printf("the total was %+v\n", total)
	fmt.Printf("the average is %+v\n", totavg)

	return totavg
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
