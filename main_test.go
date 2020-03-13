package main

import (
	"fmt"
	"testing"
)

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

	firstaverage := CalcAverage(tables[0].y, tables[1].x)
	fmt.Printf("firstaverage is %+v \n", firstaverage)

	secondaverage := CalcAverage(tables[0].x, tables[0].n, tables[1].y)
	fmt.Printf("secondaverage is %+v \n", secondaverage)

	thirdaverage := CalcAverage(tables[0].y, tables[1].n)
	fmt.Printf("thirdaverage is %+v \n", thirdaverage)

}
