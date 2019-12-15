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
	solution1 := count_centers("COM",0,to_satelites)
	fmt.Println("Solution 1: ", solution1)
	solution2 := distance_to_santa(to_center)
	fmt.Println("Solution 2: ", solution2)
}

func count_centers(obj string, curr int, to_satelites map[string][]string) int {
	satellites := to_satelites[obj]
	centers := curr
	for i := 0; i < len(satellites); i++ {
		centers += count_centers(satellites[i],curr+1,to_satelites)
	}
	return centers
}

func distance_to_center(obj string, to_center map[string]string) int {
	dist := 0
	for obj != "COM" {
		dist++
		obj = to_center[obj]
	}
	return dist
}

func distance_to_santa(to_center map[string]string) int {
	you := to_center["YOU"] // your planet
	santa := to_center["SAN"] // santa's planet
	your_radius := distance_to_center(you, to_center)
	santas_radius := distance_to_center(santa, to_center)
	fmt.Println("TO CENTER (you, santa) ", your_radius, santas_radius)
	distance := 0
	for your_radius > santas_radius {
		you = to_center[you]
		distance ++
		your_radius--
	}
	for santas_radius > your_radius {
		santa = to_center[santa]
		distance ++
		santas_radius--
	}
	fmt.Println("TO CENTER (you, santa) ", your_radius, santas_radius)
	for santa != you {
		santa = to_center[santa]
		you = to_center[you]
		distance += 2
	}
	return distance
}

func to_names(lines []string) []string {
	names := make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
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
