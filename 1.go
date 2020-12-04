package main

import (
	"os"
	"fmt"
	"strconv"
	"bufio"
)

type second struct {
	first, second int
}

func main() {
	f, _ := os.Open("1")
	first := make([]int,2020)
	second := make([]second,2020)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		remainder := 2020 - value
		first[value] = value
		if first[remainder] == remainder {
			fmt.Printf("1: %v * %v = %v\n\n", value, remainder, (value * remainder))
		}
		for _,v := range first[:remainder] {
			if v > 0 && v != value {
				second[value+v].first = value
				second[value+v].second = v
			}
		}
	}
	for k, v := range second {
		if v.first > 0 {
			remainder := 2020 - k
			if first[remainder] == remainder {
				fmt.Printf("2: %v * %v * %v = %v\n\n", v.first, v.second, remainder, (v.first * v.second * remainder))
			}
		}
	}
}
