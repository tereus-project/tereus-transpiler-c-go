package scope

import "github.com/tereus-project/tereus-remixer-c-go/remixer/ast"

type IScopeItem interface {
	GetName() string
	GetTranslatedName() string

	GetType() *ast.ASTType
}

type ScopeItem struct {
	Name           string
	TranslatedName string
}

func NewScopeItem(name string, translatedName string) *ScopeItem {
	return &ScopeItem{
		Name:           name,
		TranslatedName: translatedName,
	}
}

func (s *ScopeItem) GetName() string {
	return s.Name
}

func (s *ScopeItem) GetTranslatedName() string {
	if s.TranslatedName != "" {
		return s.TranslatedName
	}

	return s.Name
}

type Scope struct {
	Parent *Scope
	Items  map[string][]IScopeItem
}

func NewScope() *Scope {
	return &Scope{
		Items: make(map[string][]IScopeItem),
	}
}

func (s *Scope) Push() {
	parent := NewScope()
	parent.Parent = s.Parent
	parent.Items = s.Items

	s.Parent = parent
	s.Items = make(map[string][]IScopeItem)
}

func (s *Scope) Pop() {
	if s.Parent == nil {
		panic("cannot pop root scope")
	}

	s.Items = s.Parent.Items
	s.Parent = s.Parent.Parent
}

func (s *Scope) Add(item IScopeItem) {
	name := item.GetName()
	s.Items[name] = append([]IScopeItem{item}, s.Items[name]...)
}

func (s *Scope) Get(name string) []IScopeItem {
	if parent := s.Parent; parent != nil {
		return append(s.Items[name], parent.Get(name)...)
	}

	return s.Items[name]
}

func (s *Scope) GetFirst(name string) IScopeItem {
	if parent := s.Parent; parent != nil {
		return parent.GetFirst(name)
	}

	if items := s.Items[name]; len(items) > 0 {
		return items[0]
	}

	return nil
}
