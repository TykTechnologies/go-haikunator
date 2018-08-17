package haikunator

import "strings"

func asArray(in string) []string {
	words := strings.Split(in, "\n")
	return words
}
