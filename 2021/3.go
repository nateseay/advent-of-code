package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pt1() {
	//f, _ := os.Open("3_input_test.txt")
	f, _ := os.Open("3_input.txt")
	scanner := bufio.NewScanner(f)
	m := make(map[int]int)
	length := 0
	lines := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if length == 0 {
			length = len(line)
		}
		lines = lines + 1
		for i, val := range line {
			if val == '1' {
				m[i] = m[i] + 1
			}
		}
	}
	fmt.Println(m, length)
	gamma := make([]string, length)
	epsilon := make([]string, length)
	for i := 0; i < length; i++ {
		if m[i] < (lines / 2) {
			gamma[i] = "0"
			epsilon[i] = "1"
		} else {
			gamma[i] = "1"
			epsilon[i] = "0"
		}
	}
	g_binary := strings.Join(gamma, "")
	e_binary := strings.Join(epsilon, "")
	g, _ := strconv.ParseInt(g_binary, 2, 64)
	e, _ := strconv.ParseInt(e_binary, 2, 64)
	fmt.Println(g, e, g*e)
}

func ugh(more bool, lines []string, place int) int64 {
	if len(lines) == 1 {
		ret, _ := strconv.ParseInt(lines[0], 2, 64)
		return ret
	}
	var ones []string
	var zeroes []string
	for _, line := range lines {
		if line[place] == '1' {
			ones = append(ones, line)
		} else {
			zeroes = append(zeroes, line)
		}
	}

	if len(ones) == len(zeroes) {
		if more {
			return ugh(more, ones, place+1)
		} else {
			return ugh(more, zeroes, place+1)
		}
	} else {
		var greater []string
		var less []string
		if len(ones) > len(zeroes) {
			greater = ones
			less = zeroes
		} else {
			greater = zeroes
			less = ones
		}
		if more {
			return ugh(more, greater, place+1)
		} else {
			return ugh(more, less, place+1)
		}
	}
}

func pt2() {
	f, _ := os.Open("3_input.txt")
	// f, _ := os.Open("3_input.txt")
	scanner := bufio.NewScanner(f)
	var ones []string
	var zeroes []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "1") {
			ones = append(ones, line)
		} else {
			zeroes = append(zeroes, line)
		}
	}
	var oxygen, co2 int64
	if len(ones) > len(zeroes) {
		oxygen = ugh(true, ones, 1)
		co2 = ugh(false, zeroes, 1)
	} else {
		oxygen = ugh(true, zeroes, 1)
		co2 = ugh(false, ones, 1)
	}
	fmt.Println(oxygen, co2, oxygen*co2)
}

func main() {
	//pt1()
	pt2()
}
