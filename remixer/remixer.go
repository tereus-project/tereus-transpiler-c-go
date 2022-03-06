package remixer

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tereus-project/tereus-remixer-c-go/parser"
)

func Remix(entrypoint string) error {
	code, e := os.ReadFile(entrypoint)
	if e != nil {
		panic(e)
	}

	input := antlr.NewInputStream(string(code))
	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	return nil
}
