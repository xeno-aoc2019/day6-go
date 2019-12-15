package main

import (
	"fmt"
	_ "fmt"
	"io/ioutil"
	"strings"
)

func main() {
	lines := readLines()
	println("length: ",len(lines))

	to_center := map_to_center(lines)
	names := to_names(lines)
	to_satelites := map_to_satelites(names, lines)
	fmt.Println(len(names), len(to_center), len(to_satelites))
}

func to_names(lines []string) []string {
	names := make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
//		fmt.Println(i)
//		fmt.Printf("%v %v\n", lines[i], i)
		tokens := strings.Split(lines[i], ")")
		if len(tokens) > 1 {
			names[i] = tokens[1]
		}
	}
	names[len(lines)-1] = "COM"
	return names
}

func map_to_satelites(names []string, lines []string) map[string][]string {
	mapToSatelites := make(map[string][]string)
	for _, current := range names {
		var satellites = make([]string, 0)
		for _, line := range lines {
			tokens := strings.Split(line, ")")
			if tokens[0] == current {
				satellites = append(satellites, tokens[1])
			}
		}
		mapToSatelites[current] = satellites
	}
	return mapToSatelites
}

func map_to_center(lines []string) map[string]string {
	mapToCenter := make(map[string]string)
	for _, line := range lines {
		tokens := strings.Split(line, ")")
		if len(tokens) > 1 {
			mapToCenter[tokens[1]] = tokens[0]
		}
	}
	return mapToCenter
}

func readLines() []string {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	input_data := string(dat)
	lines := strings.Split(input_data, "\n")
//	fmt.Println(lines)
	return lines
}

func check(err error) {
	if err != nil {
		panic("Could not open file")
	}
}
