package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"slices"
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
			echo(params)
		} else if command == "type" {
			typeCommand(params)
		} else if command == "pwd" {
			pwd, _ := os.Getwd()
			fmt.Println(pwd)
		} else if command == "cd" {
			changeDirectore(params)
		} else if utils.SearchExecFile(command) != "" {
			executeCommand(command, strings.Split(params, " "))
		} else {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}

func echo(message string) {
	fmt.Println(strings.TrimPrefix(message, " "))
}

func executeCommand(command string, args []string) {
	cmd := exec.Command(command, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Command execution failed: %v\nOutput: %s", err, string(output))
	}

	fmt.Print(string(output))
}

func typeCommand(command string) {
	if slices.Contains(availableCommands, command) {
		fmt.Printf("%s is a shell builtin\n", command)
	} else {
		execPath := utils.SearchExecFile(command)
		if execPath != "" {
			fmt.Printf("%s is %s\n", command, execPath)
		} else {
			fmt.Printf("%s: not found\n", command)
		}
	}
}

func changeDirectore(params string) {
	directionPath := strings.Trim(params, " ")

	// this is needed just to verify how many paramas passed to the ccd command
	paramsSlice := strings.Split(directionPath, " ")
	if len(paramsSlice) > 1 {
		fmt.Println("cd: too many arguments")
		return
	}

	if directionPath == "" {
		directionPath = "~"
	}

	fullDirection := ""

	if strings.HasPrefix(directionPath, "~") {
		homeDir, _ := os.UserHomeDir()
		fullDirection = strings.ReplaceAll(directionPath, "~", homeDir)
	} else if strings.HasPrefix(directionPath, "/") {
		fullDirection = directionPath
	} else {
		cwd, _ := os.Getwd()
		fullDirection = path.Join(cwd, directionPath)
	}

	stats, err := os.Stat(fullDirection)

	if err != nil || !stats.IsDir() {
		fmt.Printf("cd: %s: No such file or directory\n", directionPath)
		return
	}

	os.Chdir(fullDirection)
}
