package language

import "strings"

type IDeclarator interface {
	String() string
}

type Declarator struct {
	Identifier string
}

func NewDeclarator(Identifier string) *Declarator {
	return &Declarator{
		Identifier: Identifier,
	}
}

func (d *Declarator) String() string {
	return d.Identifier
}

type DeclaratorFunction struct {
	Declarator IDeclarator
	Parameters []IDeclarator
}

func NewDeclaratorFunction(Declarator IDeclarator, Parameters []IDeclarator) *DeclaratorFunction {
	return &DeclaratorFunction{
		Declarator: Declarator,
		Parameters: Parameters,
	}
}

func (d *DeclaratorFunction) String() string {
	parameters := make([]string, len(d.Parameters))

	for i, parameter := range d.Parameters {
		parameters[i] = parameter.String()
	}

	return d.Declarator.String() + "(" + strings.Join(parameters, ", ") + ")"
}

func (d *DeclaratorFunction) Identifier() string {
	if declarator := d.Declarator.(*Declarator); declarator != nil {
		return declarator.Identifier
	}

	return ""
}

type DeclaratorVariable struct {
	Declarator IDeclarator
	Specifiers *Specifiers
}

func NewDeclaratorVariable(Identifier string, Specifiers *Specifiers) *DeclaratorVariable {
	return &DeclaratorVariable{
		Declarator: NewDeclarator(Identifier),
		Specifiers: Specifiers,
	}
}
