package ast

import (
	"fmt"
	"strings"
)

type ASTStruct struct {
	Name string

	IsOpaque   bool
	Properties []*ASTStructProperty
}

type ASTStructProperty struct {
	Name string
	Type *ASTType
}

func NewAstStructOpaque(name string) *ASTStruct {
	return &ASTStruct{
		Name:     name,
		IsOpaque: true,
	}
}

func (s *ASTStruct) SetProperties(properties []*ASTStructProperty) *ASTStruct {
	s.IsOpaque = false
	s.Properties = properties
	return s
}

func (s *ASTStruct) String() string {
	properties := make([]string, len(s.Properties))

	if !s.IsOpaque {
		for i, property := range s.Properties {
			properties[i] = "\t" + property.String()
		}
	}

	return fmt.Sprintf("type %s struct {\n%s\n}", s.Name, strings.Join(properties, "\n"))
}

func (s *ASTStruct) TypeString() string {
	return s.Name
}

func NewASTStructProperty(name string, typ *ASTType) *ASTStructProperty {
	return &ASTStructProperty{
		Name: name,
		Type: typ,
	}
}

func (m *ASTStructProperty) String() string {
	return fmt.Sprintf("%s %s", m.Name, m.Type.String())
}
