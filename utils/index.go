/* Package utils */
package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func CommandReader(commandStr string) (string, string) {
	firstSpaceIndex := strings.Index(commandStr, " ")

	var (
		command string
		params  = ""
	)

	if firstSpaceIndex != -1 {
		command = commandStr[0:firstSpaceIndex]
		params = commandStr[firstSpaceIndex+1:]
	} else {
		command = commandStr
	}

	return command, params
}

func SearchExecFile(command string) string {
	pathStr := os.Getenv("PATH")
	paths := strings.Split(pathStr, string(os.PathListSeparator))

	if strings.ContainsRune(command, os.PathSeparator) {
		if IsExecutable(command) {
			return command
		}
		return ""
	}

	for _, dir := range paths {
		if dir == "" {
			dir = "."
		}
		candidate := filepath.Join(dir, command)
		if IsExecutable(candidate) {
			return candidate
		}
	}
	return ""
}

func IsExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if !info.Mode().IsRegular() {
		return false
	}
	return info.Mode().Perm()&0o111 != 0
}

func FormatMessage(rowMsg string) string {
	quote := ""
	message := ""
	temp := ""
	leadingSpace := false

	for _, char := range rowMsg {
		if quote == "" && (char == rune('\'') || char == rune('"')) {
			quote = string(char)
			continue
		}

		if quote != "" {
			leadingSpace = false
			if string(char) == quote {
				message += temp
				temp = ""
				quote = ""
			} else {
				temp += string(char)
			}
		} else {
			if string(char) == " " {
				if leadingSpace {
					continue
				}

				message += string(char)
				leadingSpace = true
			} else if string(char) != " " {
				message += string(char)
				leadingSpace = false
			}
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
