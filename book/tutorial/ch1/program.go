package ch1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Hello() {
	fmt.Println("Hello this is First chapter of gopl book")
}

func EchoOsArgs() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}

func EchoOsArgs2() {
	s, sep := "", ""

	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func EchoOsArgs3() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

// Duplicate Program for detecting which word is repeated more than one, in duplicate 3 function it's checked for files

func Duplicate1() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		count[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Duplicate2() {
	count := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countFile(os.Stdin, count)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Println("Encounted error while opening file", err)
				continue
			}
			countFile(file, count)
			file.Close()
		}
	}

	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countFile(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		count[input.Text()]++
	}
}

type fileCountName struct {
	name  []string
	count int
}

func Duplicate3() {
	count := make(map[string]fileCountName)

	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Encountered error while reading the file", err.Error())
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			fnc := count[line]
			fnc.count++
			fnc.name = append(fnc.name, filename)
			count[line] = fnc
		}
	}

	for line, n := range count {
		fmt.Printf("%d\t%s\t%s\n", n.count, n.name, line)
	}
}

// Program for generating gif..

func GenerateGif() {

}
