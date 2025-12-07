package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type Sequence struct {
	startingNumber int
	lastNumber     int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func processInvalidIds(seq Sequence) int {
	fmt.Println(seq)
	var invalidIdSum int = 0
	for i := seq.startingNumber; i <= seq.lastNumber; i++ {
		numberAsString := strconv.Itoa(i)
		numberLength := len(numberAsString)
		if numberLength%2 != 0 {
			continue
		}
		half := numberLength / 2
		// fmt.Printf("Combo %s : %s\n", string(numberAsString[0:half]), string(numberAsString[half:]))
		if numberAsString[0:half] == numberAsString[half:] {
			invalidIdSum += i
			fmt.Printf("Number : %s, has %d digits and their halves are the same\n", numberAsString, len(numberAsString))
		}
	}

	return invalidIdSum
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()
	dashByte := byte('-')
	commaByte := byte(',')
	newLineByte := byte('\n')
	buffer := make([]byte, 1)
	numberToken := make([]byte, 0, 32)
	var invalidIdSum int = 0
	var seq Sequence
	for {
		bytesRead, err := f.Read(buffer)
		if err == io.EOF && bytesRead == 0 {
			break
		} else {
			check(err)
		}
		for i := 0; i < bytesRead; i++ {
			if dashByte == buffer[i] {
				seq.startingNumber, err = strconv.Atoi(string(numberToken))
				check(err)
				numberToken = numberToken[:0]
			} else if commaByte == buffer[i] {
				seq.lastNumber, err = strconv.Atoi(string(numberToken))
				check(err)
				numberToken = numberToken[:0]
				invalidIdSum += processInvalidIds(seq)
			} else if newLineByte != buffer[i] {
				numberToken = append(numberToken, buffer[i])
			}
		}
	}
	if len(numberToken) > 0 {
		seq.lastNumber, err = strconv.Atoi(string(numberToken))
		check(err)
		numberToken = numberToken[:0]
		invalidIdSum += processInvalidIds(seq)
	}
	check(err)
	
	fmt.Printf("Invalid id sum : %d\n", invalidIdSum)
}
