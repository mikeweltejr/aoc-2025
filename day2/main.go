package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	id_ranges := extractIdRanges(lines)

	return id_ranges
}

func extractIdRanges(lines []string) []string {
	var id_ranges []string

	for _, line := range lines {
		id_ranges = append(id_ranges, strings.Split(line, ",")[:]...)
	}

	return id_ranges
}

func checkIdLengthEqualAndOdd(a string, b string) bool {
	if len(a)%2 != 0 && len(b)%2 != 0 && len(a) == len(b) {
		return true
	}

	return false
}

func getInvalidIds(id_ranges []string) (int, int) {
	var invalid_ids []int
	var invalid_repeating_ids []int
	for _, id_range := range id_ranges {
		ids := strings.Split(id_range, "-")
		id_start := ids[0]
		id_end := ids[1]

		arr1, arr2 := findInvalidIds(id_start, id_end)
		invalid_ids = append(invalid_ids, arr1...)
		invalid_repeating_ids = append(invalid_repeating_ids, arr2...)
	}

	sum := addInvalidIds(invalid_ids)
	sum_repeating := addInvalidIds(invalid_repeating_ids)

	return sum, sum_repeating
}

func addInvalidIds(ids []int) int {
	sum := 0
	for _, i := range ids {
		sum += i
	}

	return sum
}

func findInvalidIds(id_start string, id_end string) ([]int, []int) {
	var invalid_ids []int
	var invalid_repeating_ids []int
	num_start, err_start := strconv.Atoi(id_start)
	num_end, err_end := strconv.Atoi(id_end)

	if err_start != nil || err_end != nil {
		fmt.Println("Error converting ids to num")
		return nil, nil
	}

	for i := num_start; i <= num_end; i++ {
		if isTwoHalvesEqual(i) {
			invalid_ids = append(invalid_ids, i)
		}

		if isRepeatingDigits(i) {
			invalid_repeating_ids = append(invalid_repeating_ids, i)
		}
	}

	return invalid_ids, invalid_repeating_ids
}

func isTwoHalvesEqual(n int) bool {
	s := strconv.Itoa(n)

	if len(s)%2 != 0 {
		return false
	}

	mid := len(s) / 2
	return s[:mid] == s[mid:]
}

func isRepeatingDigits(n int) bool {
	s := strconv.Itoa(n)

	if len(s) < 2 {
		return false
	}

	doubled := s + s

	return strings.Contains(doubled[1:len(doubled)-1], s)
}

func main() {
	id_ranges := readFileInput("input.txt")
	sum, sum_repeating := getInvalidIds(id_ranges)

	fmt.Printf("Sum of Invalid IDs: %d \n", sum)
	fmt.Printf("Sum of Repeating Digits: %d \n", sum_repeating)
}
