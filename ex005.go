package main

import "fmt"

func main () {
	vetor := [3]int{4, 5, 6}
	//		 [...]int{1,2,3} tambem funciona

	slice := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice = append(slice, slice2...)
	slice = append(slice, 7)

	//array não muda de tamanho
	//slice é um array melhor
	//slice pode mudar de tamanho

	for i, v := range slice {
		fmt.Printf("Posição: %v Valor: %v\n", i, v)
	}

	fmt.Printf("Tamanho: %v Capacidade %v\n", len(slice), cap(slice))
	fmt.Println(vetor)
}