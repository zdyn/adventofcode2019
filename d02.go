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
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			ic := NewIntcode(codes)
			ic.codes[1] = n
			ic.codes[2] = v
			for ic.state != DONE {
				ic.Run()
			}
			if n == 12 && v == 2 {
				fmt.Println(ic.codes[0])
			}
			if ic.codes[0] == 19690720 {
				fmt.Println(100*n + v)
			}
		}
	}
}
