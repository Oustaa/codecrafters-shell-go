/* Package utils */
package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func CommandReader(commandStr string) (string, string) {
	commandSlice := strings.Split(commandStr, " ")
	command := commandSlice[0]
	params := commandSlice[1:]

	return command, strings.Join(params, " ")
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
