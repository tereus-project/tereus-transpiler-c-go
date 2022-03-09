package ast

import (
	"fmt"
	"strings"
)

type ASTFunction struct {
	Name       string
	Args       []*ASTFunctionArgument
	ReturnType *ASTType
	Body       *ASTBlock
}

func NewASTFunction(name string) *ASTFunction {
	return &ASTFunction{
		Name: name,
	}
}

func (f *ASTFunction) String() string {
	body := ""
	if f.Body != nil {
		body = f.Body.String()
	}

	return fmt.Sprintf("%s %s", f.TypeString(true), body)
}

func (f *ASTFunction) TypeString(withName bool) string {
	name := ""
	if withName {
		name = f.Name
	}

	args := make([]string, len(f.Args))
	for i, arg := range f.Args {
		args[i] = arg.String()
	}

	returnType := ""
	if f.ReturnType != nil && !f.ReturnType.IsVoid() {
		returnType = " " + f.ReturnType.String()
	}

	return fmt.Sprintf("func %s(%s)%s", name, strings.Join(args, ", "), returnType)
}

type ASTFunctionArgument struct {
	Name string
	Type *ASTType
}

func NewASTFunctionArgument(name string, typ *ASTType) *ASTFunctionArgument {
	return &ASTFunctionArgument{
		Name: name,
		Type: typ,
	}
}

func (a *ASTFunctionArgument) String() string {
	return fmt.Sprintf("%s %s", a.Name, a.Type.String())
}
