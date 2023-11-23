package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/vitorsavian/interpreter/lexer"
	"github.com/vitorsavian/interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("%s", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			if tok.Type == token.IDENT && tok.Literal == "exit" {
				return
			}
			fmt.Printf("%+v\n", tok)
		}
	}
}
