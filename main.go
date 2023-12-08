package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	directions := scanner.Text()
	nodeMap := make(map[string][]string)
	if !scanner.Scan() {
		return
	}
	for scanner.Scan() {
		text := scanner.Text()
		nodeMap[text[0:3]] = []string{text[7:10], text[12:15]}

	}
	steps := 0
	cur := "AAA"
	size := len(directions)
	var pair []string
	for {
		pair = nodeMap[cur]
		if directions[steps%size] == 'L' {
			cur = pair[0]
		} else {
			cur = pair[1]
		}
		steps++
		if cur == "ZZZ" {
			break
		}
	}

	fmt.Println(steps)
}
