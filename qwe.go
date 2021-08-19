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
		return
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
		return
	}
	if val, ok := Map[cmd[1]]; !ok {
		fmt.Println("Key doesn't exists")
	} else {
		fmt.Println(val)
	}
}

func Delete(cmd []string){
	if len(cmd) != 2 {
		fmt.Println("Invalid Argument List")
		return
	}
	if _, ok := Map[cmd[1]]; !ok {
		fmt.Println("Key doesn't exists")
	} else {
		delete(Map, cmd[1])
	}
}

func About(cmd []string) {
	if len(cmd) != 2 {
		fmt.Println("Invalid Argument List")
		return
	}
	fmt.Println("qwe is a in-memory non persistent key value store.\nVersion : 1.0")
}

func handleCommand(command []string) {
	switch command[0] {
		case "SET":
			Set(command)
		case "GET":
			Get(command)
		case "DELETE":
			Delete(command)
		case "STOP":
			fmt.Println("Bye. Hope to see you again!")
			os.Exit(0)
		case "ABOUT":
			About(command)
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
