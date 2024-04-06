package main

import (
	"bufio"
	ans "ex00/anscombe"
	cli "ex00/cmd"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNumbers() ([]float64, error) {
	numbers := make([]float64, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			number, err := strconv.ParseFloat(text, 64)
			if err != nil {
				fmt.Printf("Parsing '%s' to float64 has failed\n", text)
				continue
			}
			numbers = append(numbers, number)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	return numbers, nil
}

func main() {
	cli.Execute()

	nums, err := ReadNumbers()
	fmt.Println() // to separate the outputs of reading and processing sections

	if err != nil {
		fmt.Fprintln(os.Stderr, "ReadNumbersError")
		os.Exit(3) // see the cmd/root.go module
	}

	result := make(map[string]float64)
	if cli.Flags.Mean {
		mean, err := ans.GetMean(nums)
		if err != nil {
			fmt.Fprintln(os.Stderr, "GetMean Error")
			os.Exit(4)
		}
		result["Mean"] = mean
	}
	if cli.Flags.Median {
		median, err := ans.GetMedian(nums)
		if err != nil {
			fmt.Fprintln(os.Stderr, "GetMedian Error")
			os.Exit(5)
		}
		result["Median"] = median
	}
	if cli.Flags.Mode {
		mode, err := ans.GetMode(nums)
		if err != nil {
			fmt.Fprintln(os.Stderr, "GetMode Error")
			os.Exit(6)
		}
		result["Mode"] = mode
	}
	if cli.Flags.Psd {
		psd, err := ans.GetPopulationStandardDeviation(nums)
		if err != nil {
			fmt.Fprintln(os.Stderr, "GetPopulationStandardDeviation Error")
			os.Exit(7)
		}
		result["SD"] = psd
	}
	if cli.Flags.Ssd {
		ssd, err := ans.GetSampleStandardDeviation(nums)
		if err != nil {
			fmt.Fprintln(os.Stderr, "GetSampleStandardDeviation Error")
			os.Exit(8)
		}
		result["SD"] = ssd
	}

	fmt.Println("Mean: ", result["Mean"])
	fmt.Println("Median: ", result["Median"])
	fmt.Println("Mode: ", result["Mode"])
	fmt.Println("SD: ", result["SD"])
}
