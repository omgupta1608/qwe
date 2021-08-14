package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func handleCommand(command []string) {
	switch command[0] {
		case "HI":
			fmt.Println("HELLO!")
		case "STOP":
			fmt.Println("Bye. Hope to see you again!")
			os.Exit(0)
		default:
			fmt.Println("Invalid command", command[0])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("=> ")
		text, _ := reader.ReadString('\n')
		command := strings.Fields(text)

		handleCommand(command)
	}
}
