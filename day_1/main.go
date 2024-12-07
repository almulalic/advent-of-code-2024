package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalculateDistance(left []int, right []int) int {
	distance := 0

	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	return distance
}

func CalculateSimilarityScore(left []int, right []int) int {
	score := 0
	tmp := 0

	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			if left[i] != right[j] {
				tmp += 1
				continue
			}
		}

		score += left[i] * tmp
		tmp = 0
	}

	return score
}

func main() {
	left, right, err := ParseInput("./input.txt")

	if err != nil {
		fmt.Println(fmt.Errorf("error during content parsing. %w", err))
		return
	}

	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	distance := CalculateDistance(left, right)
	similarityScore := CalculateSimilarityScore(left, right)

	fmt.Println("Distance: ", distance)
	fmt.Println("Similarity Score: ", similarityScore)
}

func ParseInput(path string) ([]int, []int, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v", err)
		}
	}(file)

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) >= 2 {
			leftNum, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing left number: %w", err)
			}

			rightNum, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing right number: %w", err)
			}

			left = append(left, leftNum)
			right = append(right, rightNum)
		} else {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return left, right, nil
}
