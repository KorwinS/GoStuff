package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	getParams()
}

func getParams() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter environment (Test, Beta, Prod): ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
