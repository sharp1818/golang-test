package main

import (
	"fmt"
	"strings"
)

var (
	stack     = []string{}
	packet    = []string{}
	needopen  = []string{}
	needclose = []string{}
	open      = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	close = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
)

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func IndexOf(stack []string, needle string) int {
	for i, v := range stack {
		if v == needle {
			return i
		}
	}
	return -1
}

func remove(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func result(s string) {
	array := strings.Split(s, "")
	for _, element := range array {
		if closer, ok := open[element]; ok {
			stack = append(stack, closer)
		} else if len(stack) == 0 {
			needopen = append(needopen, element)
		} else if element == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else if contains(stack, element) {
			remove(stack, IndexOf(stack, element))
			packet = append(packet, close[element])
		} else {
			needopen = append(needopen, element)
		}
	}

	for _, e := range stack {
		needclose = append(needclose, close[e])
	}

	if len(array) == 0 {
		fmt.Println("Empty String")
	} else if len(needclose) == 0 && len(packet) == 0 && len(needopen) == 0 {
		fmt.Println("Valid String")
	} else if len(needclose) > 0 && len(packet) > 0 {
		fmt.Println("Invalid String", packet[0], needclose[0])
	} else if len(needclose) > 0 {
		fmt.Println("Invalid String, grouping sign need close", needclose[0])
	} else if len(needopen) > 0 {
		fmt.Println("Invalid String, grouping sign need open", needopen[0])
	}
}

func main() {
	var test string = "[()]{}{[()()]()}"
	result(test)
}
