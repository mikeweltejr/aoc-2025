package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func readFileInput(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Print("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input, scanner.Err()
}

func convertInputToPoints() map[Point]string {
	input, err := readFileInput("input.txt")

	if err != nil {
		fmt.Println("Error reading input")
	}

	pointMap := make(map[Point]string)

	for i, line := range input {
		for j, r := range line {
			p := Point{X: j, Y: i}
			pointMap[p] = string(r)
		}
	}

	return pointMap
}

func calculateRollOfPaperAccess(points map[Point]string) {
	sum := 0
	rollCount := 0

	for p, val := range points {
		if val != "@" {
			continue
		}
		rollCount = 0

		if checkPoint(p.X-1, p.Y, points) {
			rollCount++
		}
		if checkPoint(p.X+1, p.Y, points) {
			rollCount++
		}
		if checkPoint(p.X, p.Y-1, points) {
			rollCount++
		}
		if checkPoint(p.X, p.Y+1, points) {
			rollCount++
		}
		if checkPoint(p.X-1, p.Y-1, points) {
			rollCount++
		}
		if checkPoint(p.X+1, p.Y-1, points) {
			rollCount++
		}
		if checkPoint(p.X-1, p.Y+1, points) {
			rollCount++
		}
		if checkPoint(p.X+1, p.Y+1, points) {
			rollCount++
		}

		if rollCount < 4 {
			sum++
		}
	}

	fmt.Println(sum)
}

func removeRollsOfPaper(removedRolls int, points map[Point]string) {
	totalRemoved := removedRolls

	for p, val := range points {
		if val != "@" {
			continue
		}

		rollCount := 0

		if checkPoint(p.X-1, p.Y, points) {
			rollCount++
		}
		if checkPoint(p.X+1, p.Y, points) {
			rollCount++
		}
		if checkPoint(p.X, p.Y-1, points) {
			rollCount++
		}
		if checkPoint(p.X, p.Y+1, points) {
			rollCount++
		}
		if checkPoint(p.X-1, p.Y-1, points) {
			rollCount++
		}
		if checkPoint(p.X+1, p.Y-1, points) {
			rollCount++
		}
		if checkPoint(p.X-1, p.Y+1, points) {
			rollCount++
		}
		if checkPoint(p.X+1, p.Y+1, points) {
			rollCount++
		}

		if rollCount < 4 {
			points[Point{X: p.X, Y: p.Y}] = "x"
			removedRolls++
		}
	}

	if totalRemoved != removedRolls {
		removeRollsOfPaper(removedRolls, points)
	} else {
		fmt.Println(totalRemoved)
	}
}

func checkPoint(x int, y int, points map[Point]string) bool {
	p, exists := points[Point{X: x, Y: y}]
	return exists && p == "@"
}

func main() {
	pointMap := convertInputToPoints()
	calculateRollOfPaperAccess(pointMap)
	removeRollsOfPaper(0, pointMap)
}
