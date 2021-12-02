package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World.")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputString := scanner.Text()
		fmt.Println(inputString)
	}
}
