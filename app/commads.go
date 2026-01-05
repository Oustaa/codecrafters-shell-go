package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"slices"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/utils"
)

func echo(message string) {
	fmt.Println(strings.TrimPrefix(message, " "))
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

func executeCommand(command string, args string) {
	cmd := exec.Command(command, utils.FormatMessage(args))

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
