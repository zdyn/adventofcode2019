package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("d02.txt")
	codes := make(map[int]int)
	for i, n := range strings.Split(strings.TrimSpace(string(input)), ",") {
		code, _ := strconv.Atoi(n)
		codes[i] = code
	}

	ic := NewIntcode(codes)
	ic.codes[1] = 12
	ic.codes[2] = 2
	for ic.state != DONE {
		ic.Run()
	}
	fmt.Println(ic.codes[0])

outer:
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			ic := NewIntcode(codes)
			ic.codes[1] = n
			ic.codes[2] = v
			for ic.state != DONE {
				ic.Run()
			}
			if ic.codes[0] == 19690720 {
				fmt.Println(100*n + v)
				break outer
			}
		}
	}
}
