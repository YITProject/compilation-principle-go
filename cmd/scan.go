package cmd

import (
	"fmt"
	"strings"
)

// Scan call fmt.Scan
func Scan(endRune rune) string {
	result := ""
	raw := ""
	for {
		_, err := fmt.Scan(&raw)
		if err != nil {
			return ""
		}
		idx := strings.IndexRune(raw, endRune)
		if idx > -1 {
			result += raw[:idx]
			break
		}
		result += raw
	}
	return result
}

func PrintInfo(code int, msg string) {
	if code != CodeEof {
		fmt.Printf("(%d,'%s')\n", code, msg)
	}
}

/* codes */
const (
	CodeEof = iota
	CodeBegin
	CodeIf
	CodeElse
	CodeWhile
	CodeDo
	CodeEnd
	CodeIdentifier = iota + 2
	CodeNumber
	CodeAdd
	CodeSubtract
	CodeMultiply
	CodeDivide
	CodeEqual
	CodeLess
	CodeGreater
	CodeLessGreater
	CodeLessEqual
	CodeNotEqual
	CodeEqualEqual
	CodeSemi
	CodeLeftParent
	CodeRightParent
)

var (
	KeywordMap = map[string]int{
		"begin": CodeBegin,
		"if":    CodeIf,
		"else":  CodeElse,
		"while": CodeWhile,
		"do":    CodeDo,
		"end":   CodeEnd,
	}
	SymbolMap = map[string]int{
		";": CodeSemi,
		"(": CodeLeftParent,
		")": CodeRightParent,
		">": CodeGreater,
		"<": CodeLess,
		"+": CodeAdd,
		"-": CodeSubtract,
		"*": CodeMultiply,
		"/": CodeDivide,
		"=": CodeEqual,
	}
	MultipleSymbolsMap = map[string]int{
		"<>": CodeLessGreater,
		"<=": CodeLessEqual,
		"!=": CodeNotEqual,
		"==": CodeEqualEqual,
	}
)

func ScanKeyword(search string) (int, bool) {
	key, ok := KeywordMap[search]
	if ok {
		return key, true
	}
	return 0, false
}

func ScanSymbols(search string) (int, bool) {
	if len(search) == 1 {
		code, ok := SymbolMap[search]
		if ok {
			return code, true
		}
	} else if len(search) == 2 {
		code, ok := MultipleSymbolsMap[search]
		if ok {
			return code, true
		}
	}
	return 0, false
}

type Parser struct {
	Source string
	Index  int
}

func (p *Parser) Scanner() (syn int, token string) {
	defer func() {
		PrintInfo(syn, token)
	}()

	syn = CodeEof
	token = ""

	// eof
	if p.Index >= len(p.Source) {
		return
	}
	ch := p.Source[p.Index]

	// merge " "
	for ch == ' ' {
		if p.Index == len(p.Source) {
			break
		}
		p.Index++
		ch = p.Source[p.Index]
	}

	// identifier
	if IsIdentifier(ch) {
		start := p.Index
		//token += string(ch) // use slices instead
		for p.Index < len(p.Source) {
			ch = p.Source[p.Index]
			if IsIdentifier(ch) || IsNumber(ch) {
				//token += string(ch)  // use slices instead
				p.Index++
			} else {
				break
			}
		}

		token = p.Source[start:p.Index] // use slices instead
		code, ok := ScanKeyword(token)
		if ok {
			return code, token
		}
		return CodeIdentifier, token
	}

	// number
	if IsNumber(ch) {
		start := p.Index
		//token += string(ch)  // use slices instead
		for p.Index < len(p.Source) {
			ch = p.Source[p.Index]
			if IsNumber(ch) {
				//token += string(ch)  // use slices instead
				p.Index++
			} else {
				break
			}
		}
		token = p.Source[start:p.Index] // use slices instead
		return CodeNumber, token
	}

	p.Index++

	// Symbol(s)
	searchString := string(ch)
	if p.Index < len(p.Source) {
		nextByte := p.Source[p.Index]
		if IsSymbol(nextByte) {
			searchString += string(nextByte)
			p.Index++
		}
	}

	code, ok := ScanSymbols(searchString)
	if ok {
		return code, searchString
	}
	return CodeEof, "End of"
}
