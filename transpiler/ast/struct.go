package ast

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type ASTStruct struct {
	TypeId string
	Name   string

	IsOpaque   bool
	Properties []*ASTStructProperty
}

func NewAstStructOpaque(name string) *ASTStruct {
	return &ASTStruct{
		TypeId:   uuid.New().String(),
		Name:     name,
		IsOpaque: true,
	}
}

func (s *ASTStruct) SetProperties(properties []*ASTStructProperty) *ASTStruct {
	s.IsOpaque = false
	s.Properties = properties
	return s
}

func (s *ASTStruct) GetProperty(name string) *ASTStructProperty {
	for _, property := range s.Properties {
		if property.Name == name {
			return property
		}
	}

	return nil
}

func (s *ASTStruct) GetPropertyByIndex(index int) *ASTStructProperty {
	return s.Properties[index]
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

func (s *ASTStruct) IsSameTo(other *ASTStruct) bool {
	return s.TypeId == other.TypeId
}

type ASTStructProperty struct {
	Name string
	Type *ASTType
}

func NewASTStructProperty(name string, typ *ASTType) *ASTStructProperty {
	return &ASTStructProperty{
		Name: name,
		Type: typ,
	}
}

func (p *ASTStructProperty) GetTranslatedName() string {
	return strings.ToUpper(string(p.Name[0])) + p.Name[1:]
}

func (m *ASTStructProperty) String() string {
	return fmt.Sprintf("%s %s", m.GetTranslatedName(), m.Type.String())
}
