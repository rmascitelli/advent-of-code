package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseWinningNums(winningNums string) map[int]bool {
	winningNumsMap := make(map[int]bool)
	var firstNum string
	for winningNums != "" {
		firstNum, winningNums, _ = strings.Cut(winningNums, " ")
		n, _ := strconv.Atoi(firstNum)
		if n != 0 {
			winningNumsMap[n] = true
		}
	}
	return winningNumsMap
}

func Part1(lines []string) {
	accumScore := 0
	for _, line := range lines {
		lineScore := 0
		fmt.Println(line)
		cardNum, scoreInfo, _ := strings.Cut(line, ": ")
		_, cardNum, _ = strings.Cut(cardNum, " ")

		winningNums, haveNums, _ := strings.Cut(scoreInfo, " | ")
		winningNumsMap := ParseWinningNums(winningNums)
		//fmt.Printf("Map - %+v\n", winningNumsMap)

		var firstNum string
		for haveNums != "" {
			firstNum, haveNums, _ = strings.Cut(haveNums, " ")
			n, _ := strconv.Atoi(firstNum)
			if _, ok := winningNumsMap[n]; ok {
				if n != 0 {
					if lineScore == 0 {
						lineScore++
					} else {
						lineScore = lineScore*2
					}
				}
			}
		}
		accumScore += lineScore
		fmt.Printf("Line score %d, accum score %d\n", lineScore, accumScore)
	}
	fmt.Println("Accum score", accumScore)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		row := scanner.Text()
		lines = append(lines, row)
	}

	//testLines := []string{"Card 15: 1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 |  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100"}
	//Part1(testLines)
	Part1(lines)

}
