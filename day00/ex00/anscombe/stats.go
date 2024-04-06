package anscombe

import (
	"errors"
	"math"
	"sort"
)

var zls_err_msg string = "zero-length sequence error"
var zls_err_val float64 = -1.0

func GetMean(nums []float64) (float64, error) {
	length := len(nums)
	if length == 0 {
		return zls_err_val, errors.New(zls_err_msg)
	}

	var sum float64
	for _, item := range nums {
		sum += item
	}
	return sum / float64(length), nil
}

func GetMedian(nums []float64) (float64, error) {
	length := len(nums)
	if length == 0 {
		return zls_err_val, errors.New(zls_err_msg)
	}

	seq := make([]float64, length)
	copy(seq, nums)
	sort.Float64s(nums)

	middle := int(length / 2)
	if length%2 == 0 {
		return (nums[middle-1] + nums[middle]) / 2.0, nil
	}
	return nums[middle], nil
}

func GetMode(nums []float64) (float64, error) {
	length := len(nums)
	if length == 0 {
		return zls_err_val, errors.New(zls_err_msg)
	}

	counter := make(map[float64]int)
	for _, item := range nums {
		amount_of_items, has_key := counter[item]
		if has_key {
			amount_of_items += 1
		} else {
			amount_of_items = 1
		}
		counter[item] = amount_of_items
	}
	mode := nums[0]
	max_freq := counter[mode]
	for num, freq := range counter {
		if freq > max_freq {
			mode, max_freq = num, freq
		} else if freq == max_freq && num < mode {
			mode = num
		}
	}
	return mode, nil
}

func get_squared_sums(nums []float64) (float64, error) {
	mean, err := GetMean(nums)
	if err != nil {
		return zls_err_val, err
	}

	var sum, diff float64
	for _, item := range nums {
		diff = (item - mean)
		sum += (diff * diff)
	}
	return sum, nil
}

func GetPopulationStandardDeviation(nums []float64) (float64, error) {
	sum, err := get_squared_sums(nums)
	if err != nil {
		return zls_err_val, err
	}
	return math.Sqrt(sum / float64(len(nums))), nil
}

func GetSampleStandardDeviation(nums []float64) (float64, error) {
	sum, err := get_squared_sums(nums)
	if err != nil {
		return zls_err_val, err
	}

	length := len(nums)
	if length < 2 {
		return zls_err_val, errors.New("sample SD cannot be computed for the length < 2")
	}
	return math.Sqrt(sum / float64(length-1)), nil
}
