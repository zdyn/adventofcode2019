package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("d06.txt")
	orbits := make(map[string][]string)
	satellites := make(map[string]string)
	for _, pair := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		bodies := strings.Split(pair, ")")
		if _, ok := orbits[bodies[0]]; !ok {
			orbits[bodies[0]] = make([]string, 0)
		}
		orbits[bodies[0]] = append(orbits[bodies[0]], bodies[1])
		satellites[bodies[1]] = bodies[0]
	}
	numOrbits := 0
	dist := 0
	queue := []string{"COM"}
	for len(queue) > 0 {
		numOrbits += len(queue) * dist
		dist++
		var next []string
		for _, body := range queue {
			next = append(next, orbits[body]...)
		}
		queue = next
	}
	fmt.Println(numOrbits)
	meDist := make(map[string]int)
	body := satellites["YOU"]
	dist = 0
	for body != "" {
		meDist[body] = dist
		body = satellites[body]
		dist++
	}
	body = satellites["SAN"]
	dist = 0
	for {
		if me, ok := meDist[body]; ok {
			fmt.Println(dist + me)
			break
		}
		body = satellites[body]
		dist++
	}
}
