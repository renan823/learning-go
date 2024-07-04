package main 

import "fmt"

const PI = 3.14159265359

func fatorial (n int) int {
	if n == 0 {
		return (1)
	}

	return (n * fatorial(n-1))
}

func potencia (x float64, n int) float64 {
	if n == 0 {
		return (1)
	}

	total := x
	for i := 1; i < n; i++ {
		total *= x;
	}

	return (total)
}

func serie (x float64, n int) float64 {
	valor := 0.0
	sinal := 1.0

	for i := 1; i <= n; i++ {
		if i % 2 != 0 {
			valor += sinal * (potencia(x, i) / float64(fatorial(i)))
			sinal *= -1
		}
	}

	return (valor)
}

func main () {
	fmt.Println(fatorial(5))
	fmt.Println(potencia(5, 2))
	fmt.Println(serie(5, 4))
}