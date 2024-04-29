package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse(config *Config) (map[string]int, error) {
	fd, err := os.Open(config.FilePath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	var result map[string]int = make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		name, balance := parse_line(line)

		stored_value, ok := result[name]
		if !ok {
			stored_value = 0
		}

		result[name] = stored_value + balance
	}

	return result, nil

}

func parse_line(line string) (name string, balance int) {

	result := strings.Split(line, ":")

	lname := strings.ToLower(result[0])
	balances := strings.Split(result[1], " ")

	l_balance := 0
	for _, b := range balances {
		i, _ := strconv.Atoi(b)
		l_balance = l_balance + i
	}

	return lname, l_balance

}
