package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Number struct {
	Length   int
	Spelling string
	Value    int
}

var StrToNum map[string][]Number

func InitStrToNum() {
	StrToNum = map[string][]Number{
		"o": {
			Number{
				Length:   3,
				Spelling: "one",
				Value:    1,
			},
		},
		"t": {
			Number{
				Length:   3,
				Spelling: "two",
				Value:    2,
			},
			Number{
				Length:   5,
				Spelling: "three",
				Value:    3,
			},
		},
		"f": {
			Number{
				Length:   4,
				Spelling: "four",
				Value:    4,
			},
			Number{
				Length:   4,
				Spelling: "five",
				Value:    5,
			},
		},
		"s": {
			Number{
				Length:   3,
				Spelling: "six",
				Value:    6,
			},
			Number{
				Length:   5,
				Spelling: "seven",
				Value:    7,
			},
		},
		"e": {
			Number{
				Length:   5,
				Spelling: "eight",
				Value:    8,
			},
		},
		"n": {
			Number{
				Length:   4,
				Spelling: "nine",
				Value:    9,
			},
		},
	}
}

func DetectNumWordFromStr(s string) Number {
	//fmt.Println("Analyzing", s)
	var n Number

	if len(s) < 3 {
		n.Length = -1
		return n
	}

	ch := s[0]
	if nums, ok := StrToNum[string(ch)]; ok {
		for _, num := range nums {
			if len(s) >= len(num.Spelling) && s[:num.Length] == num.Spelling {
				//fmt.Println("Found number - ", num.Spelling)
				return num
			}
		}
		n.Length = -1
		return n
	} else {
		n.Length = -1
		return n
	}
}

func ParseCoordinate(row string) int {
	firstCoord := -1
	lastIntScanned := -1
	for i, ch := range row {
		// If integer
		if n := DetectNumFromCh(ch); n != -1 {
			if firstCoord == -1 {
				firstCoord = n
			} else {
				lastIntScanned = n
			}
		}

		// If spelled word
		if num := DetectNumWordFromStr(row[i:]); num.Length != -1 {
			if firstCoord == -1 {
				firstCoord = num.Value
			} else {
				lastIntScanned = num.Value
			}
		}
	}
	if lastIntScanned == -1 {
		lastIntScanned = firstCoord
	}
	//fmt.Printf("Coord: %d%d\n", firstCoord, lastIntScanned)
	return firstCoord*10 + lastIntScanned
}

func DetectNumFromCh(ch rune) int {
	if int(ch) >= 48 && int(ch) <= 57 {
		return int(ch) - 48
	} else {
		return -1
	}
}

func main() {
	InitStrToNum()
	coordSum := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		fmt.Printf("row: %s", row)
		c := ParseCoordinate(row)
		fmt.Printf(", found: %d\n", c)
		coordSum += c
		fmt.Println("COORD SUM = ", coordSum)
		i++
	}
	fmt.Println("COORD SUM = ", coordSum)
}

// TESTS

func TestOneToNine() {
	ret := DetectNumWordFromStr("one")
	log.Println("Rcvd: %+v\n", ret)
	ret = DetectNumWordFromStr("two")
	log.Println("Rcvd: %+v\n", ret)
	ret = DetectNumWordFromStr("three")
	log.Println("Rcvd: %+v\n", ret)
	ret = DetectNumWordFromStr("four")
	log.Println("Rcvd: %+v\n", ret)
	ret = DetectNumWordFromStr("five")
	log.Println("Rcvd: %+v\n", ret)
	ret = DetectNumWordFromStr("six")
	log.Println("Rcvd: %+v\n", ret)
	ret = DetectNumWordFromStr("seven")
	log.Println("Rcvd: %+v\n", ret)
	ret = DetectNumWordFromStr("eight")
	log.Println("Rcvd: %+v\n", ret)
	ret = DetectNumWordFromStr("nine")
	log.Println("Rcvd: %+v\n", ret)
}
