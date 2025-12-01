package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFileInput(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Print("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dials []string
	for scanner.Scan() {
		dials = append(dials, scanner.Text())
	}
	return dials, scanner.Err()
}

func turnDials() (int, int) {
	dials, err := readFileInput("input.txt")

	if err != nil {
		fmt.Print("Error reading file")
		return -1, -1
	}

	zero_count := 0
	zero_passes := 0
	cur_dial := 50
	is_zero := false

	for _, dial := range dials {
		direction := dial[0]
		num, err := strconv.Atoi(dial[1:])
		if err != nil {
			fmt.Print("Error converting string to int")
			return -1, -1
		}

		wraps := num / 100
		moves := num - (wraps * 100)

		switch direction {
		case 'L':
			cur_dial = cur_dial - moves
		case 'R':
			cur_dial = cur_dial + moves
		}

		zero_passes += wraps
		if (cur_dial < 0 || cur_dial > 100) && !is_zero {
			zero_passes++
		}

		cur_dial = ((cur_dial % 100) + 100) % 100

		is_zero = cur_dial == 0

		if cur_dial == 0 {
			zero_count++
			zero_passes++
		}
	}

	return zero_count, zero_passes
}

func main() {
	zero_count, zero_total_count := turnDials()

	fmt.Printf("Zero Count: %s \n", strconv.Itoa(zero_count))
	fmt.Printf("Zero Passes: %s \n", strconv.Itoa(zero_total_count))
}
