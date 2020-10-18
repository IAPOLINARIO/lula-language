package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/IAPOLINARIO/lula-language/lexer"
	"github.com/IAPOLINARIO/lula-language/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.TERMINOU; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
