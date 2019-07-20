package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getResults(filename string, numberOfResults int, minHeap *IntHeap) {

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
		heap.Push(minHeap, number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("minimum: %d\n", (*minHeap)[0])

}

func main() {

	args := os.Args[1:]
	minHeap := &IntHeap{}

	if len(args) != 2 {
		log.Fatal("You must supply a file to process and a number of results to display.")
	}

	filename := args[0]
	numberOfResults, _ := strconv.Atoi(args[1])
	heap.Init(minHeap)

	getResults(filename, numberOfResults, minHeap)

}
