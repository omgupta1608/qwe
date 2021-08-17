package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

var (
	Map map[string]string = make(map[string]string, 0)
	version string = "v1"
)

func Set(cmd []string){
	if len(cmd) != 3 {
		fmt.Println("Invalid Argument List")
		return;
	}
	if _, ok := Map[cmd[1]]; ok {
		fmt.Println("Key already exists")
	} else {
		Map[cmd[1]] = cmd[2]
	}
}

func Get(cmd []string){
	if len(cmd) != 2 {
		fmt.Println("Invalid Argument List")
		return;
	}
	if val, ok := Map[cmd[1]]; !ok {
		fmt.Println("Key doesn't exists")
	} else {
		fmt.Println(val)
	}
}

func handleCommand(command []string) {
	switch command[0] {
		case "SET":
			Set(command)
		case "GET":
			Get(command)
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
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		command := strings.Fields(text)

		handleCommand(command)
	}
}
