package main

import (
	"fmt"
	_ "strings"
	// "github.com/codecrafters-io/shell-starter-go/utils"
)

func FormatMessage(rowMsg string) string {
	quote := ""
	message := ""
	temp := ""

	for _, char := range rowMsg {
		fmt.Printf("char= %s, quote= %s, temp=%s\n", string(char), quote, temp)
		if quote == "" && (char == rune('\'') || char == rune('"')) {
			quote = string(char)
			continue
		}

		if quote != "" {
			if string(char) == quote {
				message += temp
				temp = ""
				quote = ""
			} else {
				temp += string(char)
			}
		} else {
			message += string(char)
		}
	}

	if quote != "" {
		message += quote
	}

	if temp != "" {
		message += temp
	}

	return message
}

func main() {
	messages := []string{
		"\"hello 'there'\"",
		"hello i am oussma'",
		"\"He is''''' oussama\"",
	}

	fmt.Println(FormatMessage(messages[1]))
}
