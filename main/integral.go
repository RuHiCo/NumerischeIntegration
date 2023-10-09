package main

func rechteckregel(f func(float64) float64, untere_grenze float64, obere_grenze float64, n int) float64 {
	sum := 0.0
	h := (obere_grenze - untere_grenze) / float64(n)
	x_wert := untere_grenze + 0.5*h
	for i := 0; i < n; i++ {
		sum += f(x_wert) * h
		x_wert += h
	}
	return sum
}

func trapezregel(f func(float64) float64, untere_grenze float64, obere_grenze float64, n int) float64 {
	sum := 0.0
	h := (obere_grenze - untere_grenze) / float64(n)
	var y_wert_trapez float64
	x_wert := untere_grenze
	for i := 0; i < n; i++ {
		y_wert_trapez = (f(x_wert) + f(x_wert+h)) / 2.0
		x_wert = x_wert + h
		sum += y_wert_trapez * h
	}
	return sum
}

func simpsonregel(f func(float64) float64, untere_grenze float64, obere_grenze float64, n int) float64 {
	sum := 0.0
	h := (obere_grenze - untere_grenze) / float64(n)
	var y_wert_simson float64
	x_wert := untere_grenze
	for i := 0; i < n; i++ {
		y_wert_simson = (f(x_wert) + 4*f(x_wert+h/2) + f(x_wert+h)) / 6.0
		x_wert = x_wert + h
		sum += y_wert_simson * h
	}
	return sum
}

func simpson38regel(f func(float64) float64, untere_grenze float64, obere_grenze float64, n int) float64 {
	sum := 0.0
	h := (obere_grenze - untere_grenze) / float64(n)
	var y_wert_simson float64
	x_wert := untere_grenze
	for i := 0; i < n; i++ {
		y_wert_simson = (f(x_wert) + 3*f(x_wert+h/3) + 3*f(x_wert+h*2/3) + f(x_wert+h)) / 8.0
		x_wert = x_wert + h
		sum += y_wert_simson * h
	}
	return sum
}
