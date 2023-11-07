package main

import (
	"compilation-principle/cli"
	"compilation-principle/cmd"
	"fmt"
)

func main() {
	fmt.Println("输入源程序，以#结束:")
	src := ""
	if cli.Default {
		src = `begin x=90;if(x==y) y=x+78*y1; else y2=x/y1; end#`
	} else {
		src = cmd.Scan('#')
	}
	fmt.Print(src)
	fmt.Println("词法分析的结果是:")
	parser := &cmd.Parser{Source: src}
	syn, _ := parser.Scanner()
	//syn, token := scanner()
	for syn != 0 {
		//syn, token = scanner()
		syn, _ = parser.Scanner()
	}
}

/*
// /cmd/scan.go

var (
	src = ""
	index = 0
)

func scanner() (syn int, token string) {
	syn = 0
	token = ""
	// eof
	if index >= len(src) {
		return
	}
	ch := src[index]

	// merge " "
	for ch == ' ' {
		if index == len(src) {
			break
		}
		index++
		ch = src[index]
	}

	// identifier
	if cmd.IsIdentifier(ch) {
		start := index
		//token += string(ch) // use slices instead
		for index < len(src) {
			ch = src[index]
			if cmd.IsIdentifier(ch) || cmd.IsNumber(ch) {
				//token += string(ch)  // use slices instead
				index++
			} else {
				break
			}
		}

		token = src[start:index] // use slices instead

		return cmd.ScanKeyword(token)
	}

	// number
	if cmd.IsNumber(ch) {
		start := index
		//token += string(ch)  // use slices instead
		for index < len(src) {
			ch = src[index]
			if cmd.IsNumber(ch) {
				//token += string(ch)  // use slices instead
				index++
			} else {
				break
			}
		}
		token = src[start:index] // use slices instead
		return cmd.CodeNumber, token
	}
	index++
	searchString := string(ch)
	if index < len(src) {
		nextByte := src[index]
		if cmd.IsSymbol(nextByte) {
			searchString += string(nextByte)
		}
	}

	return cmd.ScanSymbols(searchString)
}
*/
