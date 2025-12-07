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
		// fmt.Printf("Number to be dissected: %d\n", i)
		isRepeating := false;
		for j := 1; j <= numberLength/2; j++ {
			if numberLength%j != 0 {
				continue
			}
			// fmt.Printf("[0,%v]", j)
			digitSequence := numberAsString[0:j]
			// seqLength := len(digitSequence)
			runesLeft := numberLength - j
			z := 0
			for k := 0; z < runesLeft/j; k = k + j {
				// fmt.Printf("Comparing %s and %s", digitSequence, numberAsString[k+j:k+j+j])
				if digitSequence != numberAsString[k+j:k+j+j] {
					// println("-nope")
					isRepeating = false
					break
				} else {
					// println("-yup")
				}
				isRepeating = true
				z++
			}
			if isRepeating {
				break;
			}
		}
		if isRepeating {
			invalidIdSum += i
			fmt.Printf("is Repeating : %d\n", i)
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
