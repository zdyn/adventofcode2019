package main

import (
	"fmt"
	"math"
)

type IntcodeState int

const (
	IDLE IntcodeState = iota
	RUNNING
	INPUT
	OUTPUT
	DONE
)

type Intcode struct {
	state  IntcodeState
	codes  map[int]int
	pos    int
	base   int
	output int
}

func NewIntcode(codes map[int]int) *Intcode {
	ic := &Intcode{codes: make(map[int]int)}
	for k, v := range codes {
		ic.codes[k] = v
	}
	return ic
}

func (ic *Intcode) Index(arg int) int {
	mode := ic.codes[ic.pos] / 100 / int(math.Pow(10, float64(arg-1))) % 10
	switch mode {
	case 0: // Position
		return ic.codes[ic.pos+arg]
	case 1: // Immediate
		return ic.pos + arg
	case 2: // Relative
		return ic.codes[ic.pos+arg] + ic.base
	}
	panic(fmt.Sprintf("unknown mode: %d", mode))
}

func (ic *Intcode) Input(val int) {
	ic.codes[ic.Index(1)] = val
	ic.pos += 2
}

func (ic *Intcode) Run() {
	ic.state = RUNNING
	code := ic.codes[ic.pos] % 100
	switch code {
	case 1: // Add
		ic.codes[ic.Index(3)] = ic.codes[ic.Index(1)] + ic.codes[ic.Index(2)]
		ic.pos += 4
		ic.Run()
		return
	case 2: // Multiply
		ic.codes[ic.Index(3)] = ic.codes[ic.Index(1)] * ic.codes[ic.Index(2)]
		ic.pos += 4
		ic.Run()
		return
	case 3: // Take input
		ic.state = INPUT
		return
	case 4: // Write output
		ic.state = OUTPUT
		ic.output = ic.codes[ic.Index(1)]
		ic.pos += 2
		return
	case 5: // Jump if true
		if ic.codes[ic.Index(1)] == 0 {
			ic.pos += 3
		} else {
			ic.pos = ic.codes[ic.Index(2)]
		}
		ic.Run()
		return
	case 6: // Jump if false
		if ic.codes[ic.Index(1)] != 0 {
			ic.pos += 3
		} else {
			ic.pos = ic.codes[ic.Index(2)]
		}
		ic.Run()
		return
	case 7: // 1 if less than, else 0
		if ic.codes[ic.Index(1)] < ic.codes[ic.Index(2)] {
			ic.codes[ic.Index(3)] = 1
		} else {
			ic.codes[ic.Index(3)] = 0
		}
		ic.pos += 4
		ic.Run()
		return
	case 8: // 1 if equal to, else 0
		if ic.codes[ic.Index(1)] == ic.codes[ic.Index(2)] {
			ic.codes[ic.Index(3)] = 1
		} else {
			ic.codes[ic.Index(3)] = 0
		}
		ic.pos += 4
		ic.Run()
		return
	case 9: // Adjust relative base
		ic.base += ic.codes[ic.Index(1)]
		ic.pos += 2
		ic.Run()
		return
	case 99: // Terminate
		ic.state = DONE
		return
	}
	panic(fmt.Sprintf("unknown code: %d", code))
}
