package main

import (
	"fmt"
)

type DataProcessor struct{}

func (dp DataProcessor) FilterData(data []int, filterFunc func(int) bool) []int {
	var result []int
	for _, v := range data {
		if filterFunc(v) {
			result = append(result, v)
		}
	}
	return result
}

func (dp DataProcessor) TransformData(data []int, transformFunc func(int) int) []int {
	var result []int
	for _, v := range data {
		result = append(result, transformFunc(v))
	}
	return result
}

func (dp DataProcessor) SumData(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}

func (dp DataProcessor) CalculateAverage(data []int) float64 {
	if len(data) == 0 {
		return 0
	}
	return float64(dp.SumData(data)) / float64(len(data))
}

// FindMax returns the maximum value in the list
func (dp DataProcessor) FindMax(data []int) int {
	if len(data) == 0 {
		panic("data slice is empty")
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func (dp DataProcessor) FindMin(data []int) int {
	if len(data) == 0 {
		panic("data slice is empty")
	}
	min := data[0]
	for _, v := range data[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	dp := DataProcessor{}

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	filteredData := dp.FilterData(data, func(x int) bool { return x%2 == 0 })
	fmt.Println("Filtered Data:", filteredData) // Output: [2 4 6 8]

	transformedData := dp.TransformData([]int{1, 2, 3, 4, 5}, func(x int) int { return x * 2 })
	fmt.Println("Transformed Data:", transformedData) // Output: [2 4 6 8 10]

	numericData := []int{10, 20, 30, 40, 50}
	totalSum := dp.SumData(numericData)
	average := dp.CalculateAverage(numericData)
	fmt.Println("Total Sum:", totalSum) // Output: 150
	fmt.Println("Average:", average)    // Output: 30

	numericData = []int{15, 5, 35, 25, 45}
	maxValue := dp.FindMax(numericData)
	minValue := dp.FindMin(numericData)
	fmt.Println("Maximum Value:", maxValue) // Output: 45
	fmt.Println("Minimum Value:", minValue) // Output: 5
}
