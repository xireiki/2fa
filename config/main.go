package config

import (
	"os"
	"strconv"
	"encoding/csv"
)

type ConfigCSV struct {
	Name    string
	Issuer  string
	Secret  string
	Digits  int
	Period  int
	Counter int
	Type    string
}

func (c *ConfigCSV) ReadConfig(filePath string) (Config []ConfigCSV, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var configOption []ConfigCSV
	for n, line := range lines {
		if n == 0 {
			continue
		}
		digits, err := strconv.Atoi(line[3])
		if err != nil {
			return nil, err
		}
		period, err := strconv.Atoi(line[4])
		if err != nil {
			return nil, err
		}
		counter, err := strconv.Atoi(line[5])
		if err != nil {
			return nil, err
		}
		options := ConfigCSV {
			Name:    line[0],
			Issuer:  line[1],
			Secret:  line[2],
			Digits:  digits,
			Period:  period,
			Counter: counter,
			Type:    line[6],
		}
		configOption = append(configOption, options)
	}
	return configOption, nil
}