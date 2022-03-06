package remixer

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tereus-project/tereus-remixer-c-go/parser"
)

func Remix(entrypoint string) error {
	visitor, e := NewVisitor(entrypoint)
	if e != nil {
		return e
	}

	input := antlr.NewInputStream(visitor.Code)
	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Translation()

	output, e := visitor.VisitTranslation(tree.(*parser.TranslationContext))
	if e != nil {
		return e
	}

	fmt.Println(output)

	return nil
}
