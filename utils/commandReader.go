/* Package utils */
package utils

import "strings"

func CommandReader(commandStr string) (string, string) {
	commandSlice := strings.Split(commandStr, " ")
	command := commandSlice[0]
	params := commandSlice[1:]

	return command, strings.Join(params, " ")
}
