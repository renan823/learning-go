package main 

import "fmt"

type Pessoa struct {
	nome string
	sobrenome string
	idade byte
	profissao string
}

type Banda struct {
	nome string
	albums []string
	membros []Pessoa
}

func main () {
	chris := Pessoa{ "Chris", "Martin", 40, "Cantor/Compositor" }
	jhony := Pessoa{ "Jhony", "Buckland", 41, "Guitarrista" }
	guy := Pessoa{ "Guy", "Berriman", 40, "Baixista" }
	will := Pessoa{ "Will", "Champion", 41, "Baterista" }

	membros := []Pessoa{chris, jhony, will, guy}

	coldplay := Banda{}
	coldplay.nome = "Coldplay"
	coldplay.albums = []string{"Parachutes", "A Head Full of Dreams"}
	coldplay.membros = membros

	fmt.Println(coldplay.membros[0].nome)
}