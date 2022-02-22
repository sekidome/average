package utils

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloat(filename string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(filename)
	// terminates the program if there is no file with the name specified within the brackets above
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
