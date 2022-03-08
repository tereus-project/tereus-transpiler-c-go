package ast

import (
	"fmt"
	"strings"
)

type ASTStruct struct {
	Name       string
	Properties []*ASTStructProperty
}

type ASTStructProperty struct {
	Name string
	Type *ASTType
}

func NewASTStruct(name string, properties []*ASTStructProperty) *ASTStruct {
	return &ASTStruct{
		Name:       name,
		Properties: properties,
	}
}

func (s *ASTStruct) String() string {
	properties := make([]string, len(s.Properties))

	for i, property := range s.Properties {
		properties[i] = "\t" + property.String()
	}

	return fmt.Sprintf("type %s struct {\n%s\n}", s.Name, strings.Join(properties, "\n"))
}

func NewASTStructProperty(name string, typ *ASTType) *ASTStructProperty {
	return &ASTStructProperty{
		Name: name,
		Type: typ,
	}
}

func (m *ASTStructProperty) String() string {
	return fmt.Sprintf("%s %s", m.Name, m.Type)
}
