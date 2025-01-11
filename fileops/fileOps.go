package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	ValueText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(ValueText), 0644)
	if err != nil {
		return
	}
}

func GetFloatFromFile(FileName string) (float64, error) {
	data, err := os.ReadFile(FileName)

	if err != nil {
		return 0, errors.New("failed to find File")
	}

	ValueText := string(data)
	balance, err := strconv.ParseFloat(ValueText, 64)
	if err != nil {
		return 0, errors.New("input invalid")
	}
	return balance, err
}
