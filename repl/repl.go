package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/alfredosa/go-interpreter/lexer"
	"github.com/alfredosa/go-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		// Scan() returns true if there is a token
		// and false if there is an error or EOF
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		// Print all the tokens
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
