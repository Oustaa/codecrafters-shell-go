package main

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/utils"
)

func main() {
	command := "echo hello there"

	command, params := utils.CommandReader(command)

	fmt.Println(command)
	fmt.Println(params)
}
