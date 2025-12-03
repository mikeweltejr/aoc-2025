package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFileInput(filename string) []string {
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

	return lines
}

func traverseBatteries(batteries []string) (int, int) {
	sum := 0
	sum2 := 0
	for _, battery := range batteries {
		sum += findLargestNDigits(battery, 2)
		sum2 += findLargestNDigits(battery, 12)
	}

	return sum, sum2
}

// Used for part 1, did not scale well
func findLargestJoltage(battery string) int {
	largestFirst := 0
	largestLast := 0
	for i := 0; i < len(battery); i++ {
		digit := int(battery[i] - '0')

		if digit > largestFirst && i != len(battery)-1 {
			largestFirst = digit
			largestLast = 1
			continue
		}
		if digit > largestLast && i != 0 {
			largestLast = digit
		}
	}

	return largestFirst*10 + largestLast
}

func findLargestNDigits(s string, N int) int {
	if N <= 0 || N > len(s) {
		fmt.Println("Invalid input length to find largest digits")
		return -1
	}

	start := 0
	result := 0

	for digitsLeft := N; digitsLeft > 0; digitsLeft-- {
		maxStart := len(s) - digitsLeft
		maxIdx := start
		largestDigit := s[start]

		for i := start + 1; i <= maxStart; i++ {
			if s[i] > largestDigit {
				largestDigit = s[i]
				maxIdx = i
			}

			if largestDigit == '9' {
				break
			}
		}

		result = result*10 + int(largestDigit-'0')
		start = maxIdx + 1
	}

	return result
}

func main() {
	batteries := readFileInput("input.txt")
	sum, sum2 := traverseBatteries(batteries)

	fmt.Printf("Battery Joltage Sum: %d \n", sum)
	fmt.Printf("Battery Joltage Part 2: %d \n", sum2)
}
