package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("d05.txt")
	codes := make(map[int]int)
	for i, n := range strings.Split(strings.TrimSpace(string(input)), ",") {
		code, _ := strconv.Atoi(n)
		codes[i] = code
	}

	ic := NewIntcode(codes)
	for ic.state != DONE {
		ic.Run()
		switch ic.state {
		case INPUT:
			ic.Input(1)
		case OUTPUT:
			fmt.Println(ic.output)
		}
	}

	ic = NewIntcode(codes)
	for ic.state != DONE {
		ic.Run()
		switch ic.state {
		case INPUT:
			ic.Input(5)
		case OUTPUT:
			fmt.Println(ic.output)
		}
	}
}
