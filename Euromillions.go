package main

import (
	"math/rand"
	"sort"
)

func GenerateLines(count int) [][]int {
	lines := make([][]int, count)

	for i := 0; i < count; i++ {
		lines[i] = generateNumbers(5)
		lines[i] = append(lines[i], generateStars(2)...)
	}

	return lines
}

func generateNumbers(count int) []int {
	numbers := make([]int, 0, count)
	used := make(map[int]bool)

	for len(numbers) < count {
		num := rand.Intn(50) + 1 
		if !used[num] {
			used[num] = true
			numbers = append(numbers, num)
		}
	}

	sort.Ints(numbers)
	return numbers
}

func generateStars(count int) []int {
	stars := make([]int, 0, count)
	used := make(map[int]bool)

	for len(stars) < count {
		star := rand.Intn(12) + 1
		if !used[star] {
			used[star] = true
			stars = append(stars, star)
		}
	}

	sort.Ints(stars)
	return stars
}

