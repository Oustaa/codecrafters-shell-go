package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	/* read users commands */
	var command string
	fmt.Scanf("%s", &command)

	fmt.Printf("%s: command not found\n", command)
}
