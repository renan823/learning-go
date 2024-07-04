package main

//import "fmt"
import "os"

func main () {
	//header := []byte("nome|sobrenome|idade")
	dados := []byte("josé|ferreira|25")

	error := os.WriteFile("dados.csv", dados, 0777)

	if error != nil {
		panic(error)
	}
}