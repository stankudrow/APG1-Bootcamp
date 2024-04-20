package reader

import (
	"bufio"
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
