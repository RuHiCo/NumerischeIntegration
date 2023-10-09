package main

import (
	"fmt"
	"math"
)

func main() {

	type type_args struct {
		function        func(float64) float64
		function_string string
		untere_grenze   float64
		obere_grenze    float64
	}

	all_args := []type_args{{func(x float64) float64 { return x*x - 5*x + 2 }, "x*x - 5*x + 2", 1, 5}, {func(x float64) float64 { return math.Sin(x) }, "sin(x)", 0, math.Pi * 2}}
	regeln := map[string]func(func(float64) float64, float64, float64, int) float64{"Rechteck": rechteckregel, "Trapez": trapezregel, "Simpson": simpsonregel}
	var erg float64
	for _, args := range all_args {
		for r_key, r := range regeln {
			for i := 10; i <= 1000; i = i * 10 {
				erg = r(args.function, args.untere_grenze, args.obere_grenze, i)
				fmt.Printf("Funktion: %#+v, Regel: %#+v, n: %d, \tErgebnis: %15.13f \n", args.function_string, r_key, i, erg)
			}
		}
	}
	//erg := simsonregel(func(x float64) float64 { return x*x - 5*x + 2 }, 1, 5, 1000)
	//fmt.Printf("Funktion: %#+v, Regel: %#+v, n: %d, Ergebnis: %10.7f \n", func(x float64) float64 { return x*x - 5*x + 2 }, "simson", 1000, erg)
}
