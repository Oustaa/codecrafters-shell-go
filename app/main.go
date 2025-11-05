package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var availableCommands []string = []string{"echo", "type", "exit"}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		command = strings.ReplaceAll(command, "\n", "")

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

		switch {
		case strings.HasPrefix(command, "echo"):
			echo(command)
		case strings.HasPrefix(command, "type"):
			typeCommand(command)
		default:
			fmt.Printf("%s: command not found\n", command)
		}

	}
}

func echo(command string) {
	/* get message */
	messageSlice := strings.Split(command, "echo ")
	/* print it */
	message := strings.Join(messageSlice, " ")
	fmt.Println(strings.TrimPrefix(message, " "))
}

func typeCommand(command string) {
	commandSlice := strings.Split(command, " ")
	if len(commandSlice) > 1 && slices.Contains(availableCommands, commandSlice[1]) {
		fmt.Printf("%s is a shell builtin", commandSlice[1])
	} else {
		fmt.Printf("%s: not found", commandSlice[1])
	}
}
