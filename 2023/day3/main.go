package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func IsInteger(ch rune) (int, bool) {
	if int(ch) >= 48 && int(ch) <= 57 {
		return int(ch) - 48, true
	} else {
		return -1, false
	}
}

func DetectSymbols(s string) bool {
	for _, ch := range s {
		if !(ch >= 48 && ch <= 57) && ch != 46 {
			return true
		}
	}
	return false
}

func DetectInt(s string) (int, int) {
	val, length := 0, 0
	var next rune
	var j int
	for i, char := range s {
		if j > 0 {
			j--
			continue
		}

		if n, ok := IsInteger(char); ok {
			intList := []int{n}
			j = 1
			next = rune(s[i+j])
			n, ok = IsInteger(next)
			for ok {
				intList = append(intList, n)
				j++
				if i+j == len(s) {
					break
				}
				next = rune(s[i+j])
				n, ok = IsInteger(next)
			}
			//fmt.Printf("Found int of length %d - %v\n", j, intList)

			for k := len(intList) - 1; k >= 0; k-- {
				x := intList[k]
				mult := int(math.Pow(10, float64(len(intList)-1-k)))
				add := x * mult
				val += add
				//fmt.Printf("Found %d, mult %d, add %d, val %d\n", x, mult, add, val)
			}
			length = len(intList)
			return val, length
		}
		break
	}
	return val, length
}

func Part1(lines []string) {
	totalSum := 0
	skip, n, rowNum := 0, 0, 0
	for _, row := range lines {
		//fmt.Println("ROWNUM =", rowNum)
		for i, _ := range row {
			if skip > 0 {
				skip--
				continue
			}
			//fmt.Printf("%c", row[i])

			if n, skip = DetectInt(row[i:]); skip > 0 {
				//fmt.Printf("\nFound int of length %d - %d\n", skip, n)
				addInt := false

				leftBound := i
				if i > 0 {
					leftBound -= 1
				}

				rightBound := i + skip
				if i+skip+1 < len(row) {
					rightBound += 1
				}

				// Check prev row
				if rowNum > 0 {
					if DetectSymbols(lines[rowNum-1][leftBound:rightBound]) {
						addInt = true
					}
				}

				// Check cur row
				if DetectSymbols(lines[rowNum][leftBound:rightBound]) {
					addInt = true
				}

				// Check next row
				if rowNum < len(lines)-1 {
					if DetectSymbols(lines[rowNum+1][leftBound:rightBound]) {
						addInt = true
					}
				}

				if addInt {
					//fmt.Println("Found symbol - adding", n)
					totalSum += n
				}

				skip--
			}
		}

		//fmt.Println("ROWSUM", totalSum)
		rowNum++
	}

	fmt.Println("Sum of valid parts =", totalSum)
}

// Detect an int
// Determine its size + "scanning boundaries"
// Scan the boundaries for symbols
// If symbol, add int to total
// Increment counter size times

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		row := scanner.Text()
		lines = append(lines, row)
	}

	Part1(lines)
}
