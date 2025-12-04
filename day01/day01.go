package main

import (
	"fmt"
	"os"
	"strconv"
)

type safe struct {
	dialNumber int
}

func (s *safe) rotateDial(right bool, amount int) int {
	var resetIfNecessary func(n int) int

	resetIfNecessary = func(n int) int {
		if n > 99 {
			return resetIfNecessary(n - 100)
		} else if n < 0 {
			return resetIfNecessary(99 + n + 1)
		}

		return n
	}

	if right {
		s.dialNumber = s.dialNumber + amount
	} else {
		s.dialNumber = s.dialNumber - amount
	}

	s.dialNumber = resetIfNecessary(s.dialNumber)

	return s.dialNumber
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("input")
	check(err)

	b1 := make([]byte, 1)
	var i int = 0
	rotationBytes := make([]byte, 4)
	christmasSafe := safe{50}
	var zeroCounter int = 0
	for {
		_, err := f.Read(b1)
		check(err)
		currentByte := b1[0]
		if currentByte != 10 {
			rotationBytes[i] = currentByte
			i++
		} else {
			right := string(rotationBytes[0]) == "R"
			amount, _ := strconv.Atoi(string(rotationBytes[1:i]))
			christmasSafe.rotateDial(right, amount)
			i = 0
			if christmasSafe.dialNumber == 0 {
				zeroCounter++
				fmt.Println(zeroCounter)
			}
			// fmt.Println(string(rotationBytes))
			// fmt.Printf("christmasSafe : %v\n", christmasSafe.dialNumber)
			rotationBytes = make([]byte, 4)
		}
	}
}
