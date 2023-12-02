package main

import (
	"bufio"
	"fmt"
	"os"
)

func AsciiCodeToInt(ascii int) int {
	return ascii - 48
}

func main() {
	//coordSum := 0
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		firstCoord := -1
		lastIntScanned := -1
		fmt.Printf("row: %s\n", row)

		for _, ch := range row {
			if int(ch) >= 48 && int(ch) <= 57 {
				if firstCoord == -1 {
					firstCoord = AsciiCodeToInt(int(ch))
					continue
				}
				lastIntScanned = AsciiCodeToInt(int(ch))
			}
		}
		fmt.Printf("Found choord : %d%d\n", firstCoord, lastIntScanned)
		break
	}
}
