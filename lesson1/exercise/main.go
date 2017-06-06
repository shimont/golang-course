package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	phonebook map[string]string
)

func save(name, phone string){
	phonebook[name]=phone
}

func get(name string) (string) {
	return phonebook[name]
}

func exit() {
	fmt.Println("Wrong command syntax")
	os.Exit(0)
}

func stringSplit(text string) (string, string, string) {
	splittedData := strings.Split(text, " ")
	if len(splittedData) == 3 {
		return splittedData[0], splittedData[1], splittedData[2]
	} else if len(splittedData) == 2 {
		return splittedData[0], splittedData[1], ""
	}
	exit()
	return "", "", ""
}

func main() {
	phonebook = make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter command: get [Name] or save [Name] [Phone]")
	for scanner.Scan() {
		text := scanner.Text()
		cmd, name, phone := stringSplit(text)
		if cmd == "get" {
			fmt.Println(get(name))
		} else if cmd == "save" {
			save(name, phone)
		} else {
			exit()
		}
	}
}
