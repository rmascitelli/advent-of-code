package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14

	RED_REGEX   = "\\d+ red"
	GREEN_REGEX = "\\d+ green"
	BLUE_REGEX  = "\\d+ blue"
)

type Regexer struct {
	Red_regex   *regexp.Regexp
	Blue_regex  *regexp.Regexp
	Green_regex *regexp.Regexp
}

func NewRegexer() Regexer {
	red_regex, err := regexp.Compile(RED_REGEX)
	blue_regex, err := regexp.Compile(BLUE_REGEX)
	green_regex, err := regexp.Compile(GREEN_REGEX)
	if err != nil {
		fmt.Println("Err creating regex:", err)
	}
	return Regexer{
		Red_regex:   red_regex,
		Blue_regex:  blue_regex,
		Green_regex: green_regex,
	}
}

var gRegexer Regexer

// 12 red cubes, 13 green cubes, and 14 blue cubes
// Determine sum of the IDs of games that are possible

// Part 2 - will take a game and determine min number of cubes required to play
func GetMinimumCubesProduct(game string) int {
	var handfull string
	var red, blue, green string
	minRed, minBlue, minGreen := 0, 0, 0
	var count int
	for game != "" {
		handfull, game, _ = strings.Cut(game, "; ")
		red = gRegexer.Red_regex.FindString(handfull)
		green = gRegexer.Green_regex.FindString(handfull)
		blue = gRegexer.Blue_regex.FindString(handfull)

		if red != "" {
			red, _, _ = strings.Cut(red, " ")
			if count, _ = strconv.Atoi(red); count > minRed {
				minRed = count
			}
		}

		if green != "" {
			green, _, _ = strings.Cut(green, " ")
			if count, _ = strconv.Atoi(green); count > minGreen {
				minGreen = count
			}
		}

		if blue != "" {
			blue, _, _ = strings.Cut(blue, " ")
			if count, _ = strconv.Atoi(blue); count > minBlue {
				minBlue = count
			}
		}
	}

	fmt.Printf("Mins - Red %d, Blue %d, Green %d\n", minRed, minBlue, minGreen)
	return minRed * minBlue * minGreen
}

func IsGamePossible(game string) bool {
	var handfull string
	var red, blue, green string
	var count int
	for game != "" {
		handfull, game, _ = strings.Cut(game, "; ")
		red = gRegexer.Red_regex.FindString(handfull)
		green = gRegexer.Green_regex.FindString(handfull)
		blue = gRegexer.Blue_regex.FindString(handfull)

		if red != "" {
			red, _, _ = strings.Cut(red, " ")
			if count, _ = strconv.Atoi(red); count > MAX_RED {
				return false
			}
		}

		if green != "" {
			green, _, _ = strings.Cut(green, " ")
			if count, _ = strconv.Atoi(green); count > MAX_GREEN {
				return false
			}
		}

		if blue != "" {
			blue, _, _ = strings.Cut(blue, " ")
			if count, _ = strconv.Atoi(blue); count > MAX_BLUE {
				return false
			}
		}
	}

	return true
}

func main() {
	gRegexer = NewRegexer()
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	possibleGidCount := 0
	sumOfMinCubePowers := 0

	for scanner.Scan() {
		row := scanner.Text()
		fmt.Println(row)

		// Get 1st game
		_, game, _ := strings.Cut(row, ": ")

		// PART 1
		//if IsGamePossible(game) {
		//	_, gidString, _ := strings.Cut(gameName, " ")
		//	gid, _ := strconv.Atoi(gidString)
		//	fmt.Println("Adding game", gid)
		//	possibleGidCount += gid
		//}

		// PART 2
		sumOfMinCubePowers += GetMinimumCubesProduct(game)

	}

	fmt.Println("Total valid game count:", possibleGidCount)
	fmt.Println("Sum of min cube powers:", sumOfMinCubePowers)
}
