package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

		fmt.Printf("%s: command not found\n", command)
	}
}
