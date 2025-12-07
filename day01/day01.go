package main

import (
	"fmt"
	"os"
	"strconv"
)

type safe struct {
	dialNumber  int
	zerosPassed int
}

func (s *safe) rotateDial(right bool, amount int) int {
	var resetIfNecessary func(n int) int

	resetIfNecessary = func(n int) int {
		if n > 99 {
			return resetIfNecessary(n - 100)
		} else if n < 0 {
			return resetIfNecessary(100 + n)
		}

		return n
	}

	println(s.dialNumber)
	println(right)
	println(amount)
	isZeroInitially := s.dialNumber == 0
	if right {

		s.dialNumber = s.dialNumber + amount
		wholes := (s.dialNumber) / 100

		s.zerosPassed = s.zerosPassed + wholes
	} else {
		s.dialNumber = s.dialNumber - amount
		if s.dialNumber < 0 {
			wholes := (s.dialNumber) / 100 * -1
			if isZeroInitially {
				s.zerosPassed = s.zerosPassed + wholes
			} else {
				s.zerosPassed = s.zerosPassed + 1 + wholes
			}
		} else if s.dialNumber == 0 {
			s.zerosPassed++
		}
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
	f, err := os.Open("input1")
	check(err)
	defer f.Close()

	b1 := make([]byte, 1)
	var i int = 0
	rotationBytes := make([]byte, 5)
	christmasSafe := safe{50, 0}
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
			fmt.Printf("%v amount %v ,\n", amount, string(rotationBytes[1:i]))
			christmasSafe.rotateDial(right, amount)
			i = 0
			fmt.Printf("the dial is rotated %s to point at %v\n", string(rotationBytes), christmasSafe.dialNumber)
			fmt.Printf("christmasSafe zeros passed : %v\n", christmasSafe.zerosPassed)
			rotationBytes = make([]byte, 5)
		}
	}
}
