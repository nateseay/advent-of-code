package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("hello world")
	var sonar []int
	//sonar = append(sonar, 199, 200, 208, 210, 200, 207, 240, 269, 260, 263)
	//fmt.Println(sonar)
	f, _ := os.Open("1_input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		intVar, _ := strconv.Atoi(scanner.Text())
		sonar = append(sonar, intVar)
	}
	var prev = sonar[0] + sonar[1] + sonar[2]
	var inc = 0
	for index, _ := range sonar {
		if index < 2 {
			continue
		}
		var cur_val = sonar[index] + sonar[index-1] + sonar[index-2]
		if cur_val > prev {
			inc = inc + 1
		}
		prev = cur_val
	}
	fmt.Println("This many increases:", inc)
}
