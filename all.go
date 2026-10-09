package main

import (
	"fmt"
	"math"
	"sort"
)

// Average возвращает среднее арифметическое элементов слайса []int.

func Average(ranges []int) float64 {
	if len(ranges) == 0 {
		return 0
	}
	sum := 0
	for _, num := range ranges {
		sum += num
	}
	return float64(sum) / float64(len(ranges))

}

// Range возвращает размах числовой последовательности.

func Range(list []int) int {
	if len(list) == 0 {
		return 0
	}
	maxV := list[0]
	minV := list[0]
	for _, num := range list {
		minV = min(minV, num)
		maxV = max(maxV, num)
	}
	return maxV - minV

}

func Median(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	copySlice := make([]int, len(nums))
	copy(copySlice, nums)
	sort.Ints(copySlice)

	mid := len(copySlice) / 2

	if len(copySlice)%2 == 0 {
		return (copySlice[mid] + copySlice[mid-1]) / 2
	}
	return copySlice[mid]
}

// Mode возвращает моды числовой последовательности.
// Напишите код функции
// ...
func Mode(s []int) ([]int, int) {
	// 1. Обработка пустого слайса: возвращаем пустой список и частоту 0
	if len(s) == 0 {
		return []int{}, 1
	}

	// 2. Подсчет частот
	freq := make(map[int]int)
	for _, num := range s {
		freq[num]++
	}

	var modes []int
	maxCount := 0

	for num, count := range freq {

		if count > maxCount {

			maxCount = count
			modes = []int{num}
		} else if count == maxCount {

			modes = append(modes, num)
		}
		if maxCount <= 1 {
			return []int{}, 1
		}
	}

	sort.Ints(modes)

	return modes, maxCount
}

func main() {
	lists := [][]int{
		{},
		{57},
		{78, -7},
		{99, 200, 0},
		{4, 4, 4, 3},
		{102, -7, 44, -7, 102},
		{82, -23, 1, 5, 98, 100},
		{100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
	}
	averages := []float64{
		0, 57, 36, 100, 4, 47, 44, 39938,
	}
	ranges := []int{
		0, 0, 85, 200, 1, 109, 123, 80000,
	}

	for i, list := range lists {
		if average := math.Round(Average(list)); average != averages[i] {
			fmt.Printf("average %d: %.2f != %.2f\n", i, averages[i], average)
		}
		if r := Range(list); r != ranges[i] {
			fmt.Printf("range %d: %d != %d\n", i, ranges[i], r)
		}
	}
	fmt.Println("Тестирование завершено")
}
