package main

import (
	"fmt"
	"math"
)

func main() {
	type typeArgs struct {
		function       function
		functionString string
		untereGrenze   float64
		obereGrenze    float64
		expected       float64
	}

	allArgs := []typeArgs{
		{func(x float64) float64 { return x*x - 5*x + 2 }, "x*x - 5*x + 2", 1, 5, -32 / 3.0},
		{func(x float64) float64 { return math.Abs(x*x - 5*x + 2) }, "|x*x - 5*x + 2|", 1, 5, 11.515},
		{func(x float64) float64 { return math.Sin(x) }, "sin(x)", 0, math.Pi * 2, 0},
		{func(x float64) float64 { return math.Abs(math.Sin(x)) }, "|sin(x)|", 0, math.Pi * 2, 4.0},
		{func(x float64) float64 { return math.Log(x) }, "ln(x)", 1, 5, 4.047189562170501873003797},
	}
	regeln := map[string]func(function, float64, float64, int) (float64, int64){
		"Rechteck":  rechteckregel,
		"Trapez":    trapezregel,
		"Simpson":   simpsonregel,
		"Simpson38": simpson38regel,
	}

	var erg float64
	var t int64
	for _, args := range allArgs {
		for rKey, r := range regeln {
			for i := 10; i <= 1000; i = i * 10 {
				erg, t = r(args.function, args.untereGrenze, args.obereGrenze, i)
				fmt.Printf("Funktion: %#+v, Regel: %#+v, n: %d, \tErgebnis: %+.16f, Abweichung: %.16f, Dauer: %d ns \n", args.functionString, rKey, i, erg, math.Abs(args.expected-erg), t)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
