package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pt1() {

	f, _ := os.Open("2_input.txt")
	scanner := bufio.NewScanner(f)
	x_pos := 0
	y_pos := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		distance, _ := strconv.Atoi(s[1])
		if s[0] == "forward" {
			x_pos = x_pos + distance
		} else if s[0] == "up" {
			y_pos = y_pos + distance
		} else if s[0] == "down" {
			y_pos = y_pos - distance
		}
	}
	if y_pos < 0 {
		y_pos = -y_pos
	}
	fmt.Println(x_pos, y_pos, "multiple: ", x_pos*y_pos)
}

func pt2() {
	//f, _ := os.Open("2_input_test.txt")
	f, _ := os.Open("2_input.txt")
	scanner := bufio.NewScanner(f)
	x_pos := 0
	y_pos := 0
	aim := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		change, _ := strconv.Atoi(s[1])
		if s[0] == "forward" {
			x_pos = x_pos + change
			y_pos = y_pos + (change * aim)
		} else if s[0] == "up" {
			aim = aim + change
		} else if s[0] == "down" {
			aim = aim - change
		}
	}
	if y_pos < 0 {
		y_pos = -y_pos
	}
	fmt.Println(x_pos, y_pos, "multiple: ", x_pos*y_pos)
}

func main() {
	//pt1()
	pt2()
}
