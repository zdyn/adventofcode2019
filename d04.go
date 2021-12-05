package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func valid(p int, p1 bool) bool {
	digits := make(map[int]int)
	lastDigit := 10
	for p != 0 {
		digit := p % 10
		if digit > lastDigit {
			return false
		}
		p /= 10
		digits[digit]++
		lastDigit = digit
	}
	for _, count := range digits {
		if (p1 && count >= 2) || (!p1 && count == 2) {
			return true
		}
	}
	return false
}

func main() {
	input, _ := ioutil.ReadFile("d04.txt")
	bounds := strings.Split(strings.TrimSpace(string(input)), "-")
	low, _ := strconv.Atoi(bounds[0])
	high, _ := strconv.Atoi(bounds[1])

	valid1 := 0
	valid2 := 0
	for p := low; p <= high; p++ {
		if valid(p, true) {
			valid1++
		}
		if valid(p, false) {
			valid2++
		}
	}
	fmt.Println(valid1)
	fmt.Println(valid2)
}
