package remixer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tereus-project/tereus-remixer-c-go/parser"
)

func Remix(entrypoint string) (string, error) {
	preprocessor, err := NewPreprocessor(entrypoint)
	if err != nil {
		return "", err
	}

	preprocessed, err := preprocessor.Preprocess()
	if err != nil {
		return "", err
	}

	// return preprocessed, nil

	visitor := NewVisitor(entrypoint, preprocessed)

	input := antlr.NewInputStream(visitor.Code)
	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCParser(stream)
	p.Interpreter.SetPredictionMode(antlr.PredictionModeSLL)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Translation()

	output, err := visitor.VisitTranslation(tree.(*parser.TranslationContext))
	if err != nil {
		return "", err
	}

	return output, nil
}
