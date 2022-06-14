package main

import (
	"github.com/tereus-project/tereus-transpiler-c-go/transpiler"
	"github.com/tereus-project/tereus-transpiler-std/core"
)

func main() {
	core.InitTranspiler(&core.TranspilerContextConfig{
		SourceLanguage:              "c",
		SourceLanguageFileExtension: ".c",
		TargetLanguage:              "go",
		TargetLanguageFileExtension: ".go",
		TranspileFunction:           transpiler.Remix,
	})
}
