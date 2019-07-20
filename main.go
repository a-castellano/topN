package main

import (
	"bufio"
	"fmt"
	"github.com/a-castellano/topN/top"
	"log"
	"os"
	"strconv"
)

func processFile(filename string, numberOfResults int, maxHeap *IntHeap) {

	fmt.Println(filename)
	fmt.Println(numberOfResults)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		heap.Push(maxHeap, number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("maximum: %d\n", (*maxHeap)[0])

}

func main() {

	args := os.Args[1:]
	top := TopN

	if len(args) != 2 {
		log.Fatal("You must supply a file to process and a number of results to display.")
	}

	filename := args[0]
	numberOfResults, _ := strconv.Atoi(args[1])
	fmt.Println(filename)
	fmt.Println(numberOfResults)
	fmt.Println(top.Size)

}
