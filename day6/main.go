package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFileInput(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Print("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func convertInputToStringArray(lines []string) [][]string {
	retArray := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		s := strings.Fields(l)

		retArray[i] = s
	}
	return retArray
}

func calculateSum(matrixArr [][]string) {
	xLen := len(matrixArr[0])
	yLen := len(matrixArr)
	opRow := yLen - 1

	sum := 0

	for x := 0; x < xLen; x++ {
		op := matrixArr[opRow][x]

		var local int
		if op == "*" {
			local = 1
		} else {
			local = 0
		}

		for y := 0; y < opRow; y++ {
			n, _ := strconv.Atoi(matrixArr[y][x])

			switch op {
			case "+":
				local += n
			case "*":
				local *= n
			default:
				panic(fmt.Errorf("unknown operator %q at col %d", op, x))
			}
		}

		sum += local
	}

	fmt.Println(sum)
}

func main() {
	lines, _ := readFileInput("input.txt")
	matrixArr := convertInputToStringArray(lines)
	calculateSum(matrixArr)
}
