package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main () {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}

	defer file.Close() //close file after all done

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}
}
