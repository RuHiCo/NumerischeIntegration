package main

import "time"

type function func(float64) float64

func rechteckregel(f function, untereGrenze float64, obereGrenze float64, n int) (float64, int64) {
	start := time.Now()
	sum := 0.0
	h := (obereGrenze - untereGrenze) / float64(n)
	xWert := untereGrenze + 0.5*h
	for i := 0; i < n; i++ {
		sum += f(xWert) * h
		xWert += h
	}
	t := time.Since(start)
	return sum, t.Nanoseconds()
}

func trapezregel(f function, untereGrenze float64, obereGrenze float64, n int) (float64, int64) {
	start := time.Now()
	sum := 0.0
	h := (obereGrenze - untereGrenze) / float64(n)
	var yWertTrapez float64
	xWert := untereGrenze
	for i := 0; i < n; i++ {
		yWertTrapez = (f(xWert) + f(xWert+h)) / 2.0
		xWert = xWert + h
		sum += yWertTrapez * h
	}
	t := time.Since(start)
	return sum, t.Nanoseconds()
}

func simpsonregel(f function, untereGrenze float64, obereGrenze float64, n int) (float64, int64) {
	start := time.Now()
	sum := 0.0
	h := (obereGrenze - untereGrenze) / float64(n)
	var yWertSimpson float64
	xWert := untereGrenze
	for i := 0; i < n; i++ {
		yWertSimpson = (f(xWert) + 4*f(xWert+h/2) + f(xWert+h)) / 6.0
		xWert = xWert + h
		sum += yWertSimpson * h
	}
	t := time.Since(start)
	return sum, t.Nanoseconds()
}

func simpson38regel(f function, untereGrenze float64, obereGrenze float64, n int) (float64, int64) {
	start := time.Now()
	sum := 0.0
	h := (obereGrenze - untereGrenze) / float64(n)
	var yWertSimpson float64
	xWert := untereGrenze
	for i := 0; i < n; i++ {
		yWertSimpson = (f(xWert) + 3*f(xWert+h/3) + 3*f(xWert+h*2/3) + f(xWert+h)) / 8.0
		xWert = xWert + h
		sum += yWertSimpson * h
	}
	t := time.Since(start)
	return sum, t.Nanoseconds()
}
