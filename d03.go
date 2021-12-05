package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func parseSegment(segment string) (x, y, dist int) {
	dist, _ = strconv.Atoi(segment[1:])
	switch segment[0:1] {
	case "R":
		x = 1
	case "L":
		x = -1
	case "U":
		y = 1
	case "D":
		y = -1
	}
	return
}

func main() {
	input, _ := ioutil.ReadFile("d03.txt")
	var wires [][]string
	for i, wire := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		wires = append(wires, make([]string, 0))
		for _, segment := range strings.Split(wire, ",") {
			wires[i] = append(wires[i], segment)
		}
	}
	points := make(map[Point]int)
	minDist := 0
	minSteps := 0
	for i, wire := range wires {
		point := Point{}
		steps := 0
		for _, segment := range wire {
			x, y, dist := parseSegment(segment)
			for j := 0; j < dist; j++ {
				point.x += x
				point.y += y
				steps++
				if s, ok := points[point]; ok == (i == 1) {
					if i == 0 {
						points[point] = steps
					} else {
						man := int(math.Abs(float64(point.x)) + math.Abs(float64(point.y)))
						if man < minDist || minDist == 0 {
							minDist = man
						}
						totalSteps := s + steps
						if totalSteps < minSteps || minSteps == 0 {
							minSteps = totalSteps
						}
					}
				}
			}
		}
	}
	fmt.Println(minDist)
	fmt.Println(minSteps)
}
