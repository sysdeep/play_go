package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(config *Config) {
	fd, err := os.Open(config.FilePath)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		name, balance := parse_line(line)

		fmt.Println(name, balance)
	}

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
