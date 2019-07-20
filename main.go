package main

import (
	"fmt"
	"github.com/a-castellano/topN/top"
	"log"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		log.Fatal("You must supply a file to process and a number of results to display.")
	}

	filename := args[0]
	numberOfResults, _ := strconv.Atoi(args[1])

	top, _ := top.NewTop(numberOfResults)
	fmt.Println(filename)
	fmt.Println(numberOfResults)
	fmt.Println(top.Size)

	fmt.Println(top.ShowMax())
	top.Push(3)
	top.Push(2)
	top.Push(7)
	fmt.Println(top.ShowMax())
}
