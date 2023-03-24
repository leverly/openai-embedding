package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	library := newLibrary()
	err := library.Init("filename.txt")
	if err != nil {
		fmt.Println("init failed:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\nInput:")
		scanner.Scan()
		query := scanner.Text()
		if query == "exit" {
			return
		}
		// find the related blocks
		err, value, result := library.FindSimilarBlock(query)
		if err != nil {
			fmt.Println("Find similar error:", err)
			continue
		}
		fmt.Println("find similar block:", value)
		fmt.Println(result)
	}
}
