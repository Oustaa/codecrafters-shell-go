package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/utils"
)

var availableCommands []string = []string{"echo", "type", "exit", "pwd"}

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
		} else if searchExecFile(command) != "" {
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
		execPath := searchExecFile(command)
		if execPath != "" {
			fmt.Printf("%s is %s\n", command, execPath)
		} else {
			fmt.Printf("%s: not found\n", command)
		}
	}
}

func searchExecFile(command string) string {
	pathStr := os.Getenv("PATH")
	paths := strings.Split(pathStr, string(os.PathListSeparator))

	if strings.ContainsRune(command, os.PathSeparator) {
		if isExecutable(command) {
			return command
		}
		return ""
	}

	for _, dir := range paths {
		if dir == "" {
			dir = "."
		}
		candidate := filepath.Join(dir, command)
		if isExecutable(candidate) {
			return candidate
		}
	}
	return ""
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if !info.Mode().IsRegular() {
		return false
	}
	return info.Mode().Perm()&0o111 != 0
}
