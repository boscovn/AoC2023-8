package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func findLCD(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	directions := scanner.Text()
	nodeMap := make(map[string][]string)
	var cur []string
	if !scanner.Scan() {
		return
	}
	for scanner.Scan() {
		text := scanner.Text()
		key := text[0:3]
		nodeMap[key] = []string{text[7:10], text[12:15]}
		if key[2] == 'A' {
			cur = append(cur, key)
		}

	}
	size := len(directions)
	var stepsNeeded []int
	for _, v := range cur {
		steps := 0
		for {
			pair := nodeMap[v]
			if directions[steps%size] == 'L' {
				v = pair[0]
			} else {
				v = pair[1]
			}
			steps++
			if v[2] == 'Z' {
				stepsNeeded = append(stepsNeeded, steps)
				break
			}
		}
	}

	fmt.Println(findLCD(stepsNeeded))
}
