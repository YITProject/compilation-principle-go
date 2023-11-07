package main

import (
	"compilation-principle/cli"
	"compilation-principle/cmd"
	"errors"
	"fmt"
)

var (
	syn        = 0
	parser     *cmd.Parser
	ErrGrammar = errors.New("error grammar")
)

func main() {
	fmt.Println("输入源程序，以#结束:")
	src := ""
	if cli.Default {
		src = `begin x=7;if(x==y) x=1 else y=9+l end#`
	} else {
		src = cmd.Scan('#')
	}
	fmt.Print(src)
	fmt.Println("词法分析的结果是:")
	parser = &cmd.Parser{Source: src}
	ScanNext()
	RP()
	ScanNext()
	if syn == cmd.CodeEof {
		fmt.Println("分析成功")
	} else {
		panic(ErrGrammar)
	}
}

// RP -> begin statement{;statement} end
func RP() {
	// · begin ...
	if syn == cmd.CodeBegin {
		PrintTable("Begin")
		{
			ScanNext()

			// begin · statement ...
			Statement()

			// ... { · ;statement} ...
			// · ;
			for syn == cmd.CodeSemi {
				PrintTable("Symbol ;")
				ScanNext()

				// ... {; · statement} ...
				Statement()
			}
		}
	}
	// ... · end
	if syn != cmd.CodeEnd {
		panic(ErrGrammar)
	}
}

// Statement -> id=expression | if (bool) statement [else statement]
func Statement() {
	PrintTable("Statement start")
	defer func() {
		PrintTable("Statement end")
	}()
	// · id ...
	// id
	if syn == cmd.CodeIdentifier {
		PrintTable("id")
		ScanNext()
		// id · = ...
		// =
		if syn == cmd.CodeEqual {
			PrintTable("symbol =")
			ScanNext()

			// id = · expression
			Expression()
			return
		} else {
			panic(ErrGrammar)
		}
	}
	// · if ...
	if syn == cmd.CodeIf {
		ScanNext()
		// if · (...
		if syn == cmd.CodeLeftParent {
			PrintTable("symbol (")
			ScanNext()
			// if (·bool) statement
			Bool()
			// if (bool ·) statement
			if syn == cmd.CodeRightParent {
				PrintTable("symbol )")
				ScanNext()
				// if (bool) ·statement
				Statement()
			} else {
				panic(ErrGrammar)
			}
		} else {
			panic(ErrGrammar)
		}
		// if (bool) statement · else statement
		if syn == cmd.CodeElse {
			ScanNext()
			// if (bool) statement else · statement
			Statement()
		}
		return
	}
	panic(ErrGrammar)
}

// Bool -> factor{!=factor | ==factor | >factor | <factor | <>factor | <=factor}
func Bool() {
	PrintTable("Boolean start")
	defer func() {
		PrintTable("Boolean end")
	}()
	// · factor {...}
	Factor()
	// factor {· ...factor}
	for syn == cmd.CodeNotEqual || syn == cmd.CodeEqualEqual || syn == cmd.CodeGreater || syn == cmd.CodeLess || syn == cmd.CodeLessGreater || syn == cmd.CodeLessEqual {
		PrintTable("bool symbols")
		ScanNext()
		// factor {...· factor}
		Factor()
	}
}

// Expression -> term{+term | -term}
func Expression() {
	PrintTable("Expression start")
	defer func() {
		PrintTable("Expression end")
	}()
	// · term{+term | -term}
	Term()
	// term{· +term | · -term}
	for syn == cmd.CodeAdd || syn == cmd.CodeSubtract {
		ScanNext()
		// term{+ · term | - · term}
		Term()
	}
}

// Term -> factor{*factor | /factor}
func Term() {
	PrintTable("Term start")
	defer func() {
		PrintTable("Term end")
	}()
	// · factor {*factor | /factor}
	Factor()

	// · factor {· *factor | · /factor}
	// * | /
	for syn == cmd.CodeMultiply || syn == cmd.CodeDivide {
		ScanNext()
		// · factor {*· factor | /· factor}
		Factor()
	}
}

// Factor -> id | number | (expression)
func Factor() {
	PrintTable("Factor start")
	defer func() {
		PrintTable("Factor end")
	}()
	ScanNext()
	// · (expression)
	// (
	if syn == cmd.CodeLeftParent {
		PrintTable("symbol (")
		ScanNext()

		// (· expression)
		Expression()

		// (expression · )
		// )
		if syn == cmd.CodeRightParent {
			PrintTable("symbol )")
			ScanNext()
			return
		} else {
			panic(ErrGrammar)
		}
		// · id | · number
	} else if syn == cmd.CodeIdentifier || syn == cmd.CodeNumber {
		fmt.Println("End: id or number")
	}
}

// ScanNext call parser.scanner
func ScanNext() {
	syn, _ = parser.Scanner()
}

func PrintTable(s string) {
	if cli.Table {
		fmt.Printf("     %s\n", s)
	}
}
