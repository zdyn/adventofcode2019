package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func fuelForMass(mass int) int {
	return mass/3 - 2
}

func fuelForMassAndFuel(mass int) int {
	fuel := 0
	for mass >= 9 { // 8 / 3 - 2 = 0 (rounded down)
		mass = fuelForMass(mass)
		fuel += mass
	}
	return fuel
}

func main() {
	input, _ := ioutil.ReadFile("d01.txt")
	var modules []int
	for _, n := range strings.Split(string(input), "\n") {
		if n == "" {
			continue
		}
		module, _ := strconv.Atoi(n)
		modules = append(modules, module)
	}

	fuel := 0
	for _, module := range modules {
		fuel += fuelForMass(module)
	}
	fmt.Println(fuel)

	fuel = 0
	for _, module := range modules {
		fuel += fuelForMassAndFuel(module)
	}
	fmt.Println(fuel)
}
