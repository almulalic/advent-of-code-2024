package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	levels, err := ReadAdventInput("input.txt")

	if err != nil {
		panic(err)
	}

	safeCount := 0
	dampenedSafeCount := 0

	for i := 0; i < len(levels); i++ {
		if safe, unsafeIndices := ScanReport(levels[i]); safe {
			safeCount++
			dampenedSafeCount++
		} else {
			for _, index := range unsafeIndices {
				rowCopy := make([]int, len(levels[i]))
				copy(rowCopy, levels[i])

				dampenedLevels := append(rowCopy[:index], rowCopy[index+1:]...)

				if safe := IsReportSafe(dampenedLevels); safe {
					dampenedSafeCount++
					break
				}
			}
		}
	}

	fmt.Println("Safe records: ", safeCount)
	fmt.Println("Safe records (with Problem Dampener): ", dampenedSafeCount)
}

// ScanReport (row) => (safe, unsafe indices)
func ScanReport(row []int) (bool, []int) {
	curr := 0
	prev := 0
	safe := true

	var isDescending int
	var unsafeIndices []int

	if row[0] > row[1] && row[1] > row[2] {
		isDescending = 1
	} else if row[0] < row[1] && row[1] < row[2] {
		isDescending = -1
	} else {
		isDescending = 0
	}

	for j := 1; j < len(row); j++ {
		curr = row[j]
		prev = row[j-1]
		diff := int(math.Abs(float64(curr - prev)))

		if diff <= 0 || diff > 3 {
			safe = false
			unsafeIndices = append(unsafeIndices, j, j-1)
			continue
		}

		if isDescending == 0 || isDescending == 1 && curr >= prev || isDescending == -1 && curr <= prev {
			safe = false
			unsafeIndices = append(unsafeIndices, j, j-1)
			continue
		}
	}

	return safe, unsafeIndices
}

func IsReportSafe(row []int) bool {
	curr := 0
	prev := 0
	isDescending := row[0] > row[1]

	for j := 1; j < len(row); j++ {
		curr = row[j]
		prev = row[j-1]
		diff := int(math.Abs(float64(curr - prev)))

		if diff <= 0 || diff > 3 {
			return false
		}

		if isDescending && curr >= prev || !isDescending && curr <= prev {
			return false
		}
	}

	return true
}

func ReadAdventInput(path string) ([][]int, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var levels [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		var row []int
		for _, part := range parts {
			part = strings.TrimSpace(part)
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error parsing number '%s': %v", part, err)
			}
			row = append(row, num)
		}

		levels = append(levels, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return levels, nil
}
