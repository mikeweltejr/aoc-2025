package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int64
	End   int64
}

func readFileInput(filename string) ([]string, []string, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Print("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isAvailableIngredients := false
	var freshIngredientRanges []string
	var availableIngredients []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			isAvailableIngredients = true
			continue
		}
		if !isAvailableIngredients {
			freshIngredientRanges = append(freshIngredientRanges, scanner.Text())
		} else {
			availableIngredients = append(availableIngredients, scanner.Text())
		}

	}
	return freshIngredientRanges, availableIngredients, scanner.Err()
}

func parseRanges(freshIngredients []string) []Range {
	ranges := make([]Range, 0, len(freshIngredients))
	for _, ingredientRange := range freshIngredients {
		parts := strings.Split(strings.TrimSpace(ingredientRange), "-")
		if len(parts) != 2 {
			fmt.Println("Bad file input")
			return nil
		}

		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end, _ := strconv.ParseInt(parts[1], 10, 64)
		if start > end {
			start, end = end, start
		}
		ranges = append(ranges, Range{Start: start, End: end})
	}

	return ranges
}

func inAnyRange(x int64, ranges []Range) bool {
	for _, r := range ranges {
		if x >= r.Start && x <= r.End {
			return true
		}
	}
	return false
}

func calculateAvailableFreshIngredients(freshIngredientRanges []string, availableIngredients []string) {
	ranges := parseRanges(freshIngredientRanges)

	sum := 0
	for _, ingredient := range availableIngredients {
		i, _ := strconv.ParseInt(ingredient, 10, 64)

		if inAnyRange(i, ranges) {
			sum++
		}
	}

	fmt.Println(sum)
}

func mergeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].Start == ranges[j].Start {
			return ranges[i].End < ranges[j].End
		}
		return ranges[i].Start < ranges[j].Start
	})

	merged := make([]Range, 0, len(ranges))
	cur := ranges[0]

	for _, r := range ranges[1:] {
		if r.Start <= cur.End+1 {
			if r.End > cur.End {
				cur.End = r.End
			}
		} else {
			merged = append(merged, cur)
			cur = r
		}
	}
	merged = append(merged, cur)

	return merged
}

func calculateAllFreshIngredients(freshIngredientRanges []string) {
	ranges := parseRanges(freshIngredientRanges)
	ranges = mergeRanges(ranges)

	var sum int64
	for i := 0; i < len(ranges); i++ {
		r := ranges[i]

		sum += (r.End - r.Start) + 1
	}

	fmt.Println(sum)
}

func main() {
	freshIngredientRanges, availableIngredients, _ := readFileInput("input.txt")
	calculateAvailableFreshIngredients(freshIngredientRanges, availableIngredients)
	calculateAllFreshIngredients(freshIngredientRanges)
}
