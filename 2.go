package main

import (
	"os"
	"fmt"
	"regexp"
	"bufio"
	"strconv"
)

type policy struct {
	min, max int
	char string
}

func (p *policy) isValidPasswordTypeOne(password string) bool {
	r, _ := regexp.Compile(p.char)
	matches := len(r.FindAllString(password, -1))
	if matches >= p.min && matches <= p.max {
		return true
	}
	return false
}

func (p *policy) isValidPasswordTypeTwo(password string) bool {
	matches := 0
	if password[p.min - 1] == p.char[0] { matches++ }
	if password[p.max - 1] == p.char[0] { matches++ }
	if matches == 1 {
		return true
	}
	return false
}

func main() {
	f, _ := os.Open("2")
	countTypeOne := 0
	countTypeTwo := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	r, _ := regexp.Compile(`^(\d+)-(\d+) (\w): (\w+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindStringSubmatch(line)
		if matches != nil {
			min, _ := strconv.Atoi(matches[1])
			max, _ := strconv.Atoi(matches[2])
			char := matches[3]
			policy := &policy{
				min: min,
				max: max,
				char: char,
			}
			password := matches[4]
			if policy.isValidPasswordTypeOne(password) { countTypeOne++ }
			if policy.isValidPasswordTypeTwo(password) { countTypeTwo++ }
		}
	}
	fmt.Printf("%v valid passwords with method one\n", countTypeOne)
	fmt.Printf("%v valid passwords with method two\n", countTypeTwo)
}
