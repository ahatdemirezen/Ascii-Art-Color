package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var colors = map[string]string{
	"background":    "\033[7m", // Background color
	"black":         "\033[30m",
	"red":           "\033[31m",
	"green":         "\033[32m",
	"yellow":        "\033[33m",
	"blue":          "\033[34m",
	"magenta":       "\033[35m",
	"cyan":          "\033[36m",
	"white":         "\033[37m",
	"gray":          "\033[90m",
	"brightred":     "\033[91m",
	"brightgreen":   "\033[92m",
	"brightyellow":  "\033[93m",
	"brightblue":    "\033[94m",
	"brightmagenta": "\033[95m",
	"brightcyan":    "\033[96m",
	"brightwhite":   "\033[97m",
}

func main() {
	if len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] || Example: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] || Example: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
	colorFlag := os.Args[1][:7]
	if colorFlag != "--color" {
		fmt.Println("Usage: go run . [OPTION] [STRING] || Example: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
	color := os.Args[1][8:]
	color = strings.ToLower(color)
	lettersToColor := os.Args[2]
	var argStr string
	if len(os.Args) == 4 {
		argStr = os.Args[3]
	}
	if len(os.Args) == 3 {
		argStr = os.Args[2]
	}

	sepArgs := strings.Split(argStr, "\\n")

	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}

	printColoredAsciiArt(sepArgs, lines, colors[color], lettersToColor)
}

func printColoredAsciiArt(sentences []string, textFile []string, colorCode string, paintLetters string) {
	for i, word := range sentences {
		if word == "" {
			if i != 0 { 
				fmt.Println()
			}
			continue
		}
		for h := 1; h < 9; h++ {
			for i := 0; i < len(word); i++ {
				paint := false
				for _, char := range paintLetters { 
					if string(char) == string(word[i]) {
						paint = true
						break
					}
				}
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[i])-32)*9+h {
						if paint {
							fmt.Print(colorCode + line + "\033[0m")
							break
						} else {
							fmt.Print(line)
							break
						}
					}
				}
			}
			fmt.Println()
		}
	}
}
