package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/utils"
)

var availableCommands []string = []string{"echo", "type", "exit", "pwd", "cd"}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		reader := bufio.NewReader(os.Stdin)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		userInput = strings.ReplaceAll(userInput, "\n", "")
		command, params := utils.CommandReader(userInput)

		// exit on command exit
		if strings.HasPrefix(command, "exit") {
			commandParts := strings.Split(command, " ")
			var code int64 = 0
			if len(commandParts) > 1 {
				codeStr := commandParts[1]
				code, _ = strconv.ParseInt(codeStr, 10, 8)
			}

			os.Exit(int(code))
		}

		if command == "echo" {
			echo(utils.FormatMessage(params))
		} else if command == "type" {
			typeCommand(params)
		} else if command == "pwd" {
			pwd, _ := os.Getwd()
			fmt.Println(pwd)
		} else if command == "cd" {
			changeDirectore(params)
		} else if utils.SearchExecFile(command) != "" {
			executeCommand(command, params)
		} else {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
