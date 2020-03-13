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

	var fp = flag.String("fp", "", "Pass in the filename and path \nexample ./main -fp=./input_data/input_test_data.csv (Required)")
	//var altnan = flag.String("altnan", "nan", "The default indicator of a missing number is nan but "+
	//	"you can add your own string here (Optional)")
	//var outfile = flag.String("outfile", "output.csv", "The default output file is called output.csv and is \n"+
	//	"saved in the current directory add the path and filename if you want a different name and location (Optional)")
	var help = flag.String("h", "", "Extended help")
	flag.Parse()

	if *fp == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *help == "" {

	}

	rows := readSample(*fp)
	appendSum(rows)
	writeChanges(rows)
}

func readSample(thefile string) [][]string {
	f, err := os.Open(thefile)

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
	mr := len(rows) - 1
	lr := len(rows[0]) - 1

	for cr := 0; cr <= len(rows)-1; cr++ {
		for cc := 0; cc <= len(rows[cr])-1; cc++ {
			if rows[cr][cc] == "nan" {
				if cr == 0 {
					if cc == 0 {
						a1 := rows[cr][cc+1]
						a2 := rows[cr+1][cc]
						a164, _ := strconv.ParseFloat(a1, 64)
						a264, _ := strconv.ParseFloat(a2, 64)
						newVal := CalcAverage(a164, a264)
						rows[cr][cc] = fmt.Sprintf("%f", newVal)
					} else if cc < lr {
						a1 := rows[cr][cc-1]
						a2 := rows[cr][cc+1]
						a3 := rows[cr+1][cc]
						a164, _ := strconv.ParseFloat(a1, 64)
						a264, _ := strconv.ParseFloat(a2, 64)
						a364, _ := strconv.ParseFloat(a3, 64)
						newVal := CalcAverage(a164, a264, a364)
						rows[cr][cc] = fmt.Sprintf("%f", newVal)
					} else if cc == lr {
						a1 := rows[cr][cc-1]
						a2 := rows[cr+1][cc]
						a164, _ := strconv.ParseFloat(a1, 64)
						a264, _ := strconv.ParseFloat(a2, 64)
						newVal := CalcAverage(a164, a264)
						rows[cr][cc] = fmt.Sprintf("%f", newVal)
					}
				} else if cr < mr {
					if cc == 0 {
						a1 := rows[cr][cc+1]
						a2 := rows[cr-1][cc]
						a3 := rows[cr+1][cc]
						a164, _ := strconv.ParseFloat(a1, 64)
						a264, _ := strconv.ParseFloat(a2, 64)
						a364, _ := strconv.ParseFloat(a3, 64)
						newVal := CalcAverage(a164, a264, a364)
						rows[cr][cc] = fmt.Sprintf("%f", newVal)
					} else if cc < lr {
						a1 := rows[cr][cc-1]
						a2 := rows[cr][cc+1]
						a3 := rows[cr-1][cc]
						a4 := rows[cr+1][cc]
						a164, _ := strconv.ParseFloat(a1, 64)
						a264, _ := strconv.ParseFloat(a2, 64)
						a364, _ := strconv.ParseFloat(a3, 64)
						a464, _ := strconv.ParseFloat(a4, 64)
						newVal := CalcAverage(a164, a264, a364, a464)
						rows[cr][cc] = fmt.Sprintf("%f", newVal)
					} else if cc == lr {
						a1 := rows[cr][cc-1]
						a2 := rows[cr-1][cc]
						a3 := rows[cr+1][cc]
						a164, _ := strconv.ParseFloat(a1, 64)
						a264, _ := strconv.ParseFloat(a2, 64)
						a364, _ := strconv.ParseFloat(a3, 64)
						newVal := CalcAverage(a164, a264, a364)
						rows[cr][cc] = fmt.Sprintf("%f", newVal)
					}
				} else if cr == mr {
					if cc == 0 {
						a1 := rows[cr][cc+1]
						a2 := rows[cr-1][cc]
						a164, _ := strconv.ParseFloat(a1, 64)
						a264, _ := strconv.ParseFloat(a2, 64)
						newVal := CalcAverage(a164, a264)
						rows[cr][cc] = fmt.Sprintf("%f", newVal)
					} else if cc < lr {
						a1 := rows[cr][cc-1]
						a2 := rows[cr][cc+1]
						a3 := rows[cr-1][cc]
						a164, _ := strconv.ParseFloat(a1, 64)
						a264, _ := strconv.ParseFloat(a2, 64)
						a364, _ := strconv.ParseFloat(a3, 64)
						newVal := CalcAverage(a164, a264, a364)
						rows[cr][cc] = fmt.Sprintf("%f", newVal)
					} else if cc == lr {
						a1 := rows[cr][cc-1]
						a2 := rows[cr-1][cc]
						a164, _ := strconv.ParseFloat(a1, 64)
						a264, _ := strconv.ParseFloat(a2, 64)
						newVal := CalcAverage(a164, a264)
						rows[cr][cc] = fmt.Sprintf("%f", newVal)
					}
				}
			}
		}
	}
}

func CalcAverage(nums ...float64) float64 {
	var total float64
	for _, num := range nums {
		total += num
	}
	totavg := total / float64(len(nums))
	return totavg
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
