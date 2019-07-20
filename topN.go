package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getResults(filename string, numberOfResults int) {
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
		fmt.Println(number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		log.Fatal("You must supply a file to process and a number of results to display.")
	}

	filename := args[0]
	numberOfResults, _ := strconv.Atoi(args[1])

	getResults(filename, numberOfResults)
}
