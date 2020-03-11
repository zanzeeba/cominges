package main

import (
	"fmt"
	"testing"
)

//test_data := [][]float64{
//[]float64{1.1, 1.2, 1.3},
//[]float64{2.1, 2.2, 2.3},
//[]float64{3.1, 3.2, 3.3},
//}

func TestCalcAverage(t *testing.T) {
	tables := []struct {
		x float64
		y float64
		n float64
	}{
		{1.1, 1.2, 2.3},
		{2.2, 2.2, 4.4},
		{3.2, 3.2, 6.4},
	}
	//fmt.Println(tables[0].x)
	//fmt.Println(tables[0].y)
	//fmt.Println(tables[0].n)

	firstaverage := CalcAverage(tables[0].y, tables[1].x)
	fmt.Printf("firstaverage is %+v \n", firstaverage)

	secondaverage := CalcAverage(tables[0].x, tables[0].n, tables[1].y)
	fmt.Printf("secondaverage is %+v \n", secondaverage)

	thirdaverage := CalcAverage(tables[0].y, tables[1].n)
	fmt.Printf("thirdaverage is %+v \n", thirdaverage)

	//for _, table := range tables {
	//	total := CalcAverage(table.x, table.y)
	//	if total != table.n {
	//		t.Errorf("Sum of (%f+%f) was incorrect, got: %f, want: %f.", table.x, table.y, total, table.n)
	//	}
	//}
}
