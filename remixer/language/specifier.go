package language

type Specifiers struct {
	Type *SpecifierType
}

func NewSpecifiers() *Specifiers {
	return &Specifiers{}
}

func (s *Specifiers) String() string {
	switch s.Type.Name {
	case "void":
		return ""
	case "char":
		return "Rune"
	case "short":
		return "int16"
	case "int":
		return "int"
	case "long":
		return "int64"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "signed":
		return "int"
	case "unsigned":
		return "uint"
	}

	return "???"
}

type ISpecifier interface{}

type SpecifierType struct {
	Name string
}

func NewSpecifierType(Name string) *SpecifierType {
	return &SpecifierType{
		Name: Name,
	}
}
