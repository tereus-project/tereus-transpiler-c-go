package remixer

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Define struct {
	Name       string
	IsFunction bool
	Args       []string

	Content string
}

type Preprocessor struct {
	Path string
	Code string

	Output map[string]string

	Defines map[string]*Define
}

func NewPreprocessor(path string) (*Preprocessor, error) {
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	preprocessor := &Preprocessor{
		Path:   path,
		Code:   string(code),
		Output: make(map[string]string),

		Defines: make(map[string]*Define),
	}

	return preprocessor, nil
}

var (
	preprocessorRegex         = regexp.MustCompile(`(?m)^[ \t]*#\s*define\s+(?P<Name>\w+)(?P<Args>\(\s*\w+(?:\s*,\s*\w+)*\s*\))?\s+(?P<Value>(?:(?:[\\][\r\n])|.)+)$`)
	preprocessorArgumentRegex = regexp.MustCompile(`^[a-zA-Z_]\w*$`)
)

func (v *Preprocessor) Preprocess() (string, error) {
	nameIndexStart := preprocessorRegex.SubexpIndex("Name") * 2
	nameIndexEnd := nameIndexStart + 1

	valueIndexStart := preprocessorRegex.SubexpIndex("Value") * 2
	valueIndexEnd := valueIndexStart + 1

	argsIndexStart := preprocessorRegex.SubexpIndex("Args") * 2
	argsIndexEnd := argsIndexStart + 1

	defines := make(map[string]*Define)
	definesStartIndex := make(map[string]int)

	for matches := preprocessorRegex.FindStringSubmatchIndex(v.Code); matches != nil; matches = preprocessorRegex.FindStringSubmatchIndex(v.Code) {
		name := v.Code[matches[nameIndexStart]:matches[nameIndexEnd]]
		value := v.Code[matches[valueIndexStart]:matches[valueIndexEnd]]

		if _, ok := defines[name]; ok {
			return "", fmt.Errorf("'%s' already defined", name)
		}

		if argsStart := matches[argsIndexStart]; argsStart != -1 {
			argsString := v.Code[argsStart:matches[argsIndexEnd]]
			argsString = argsString[1 : len(argsString)-1]
			args := strings.Split(argsString, ",")

			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)

				if !preprocessorArgumentRegex.MatchString(args[i]) {
					return "", fmt.Errorf("invalid argument: '%s'", args[i])
				}
			}

			defines[name] = &Define{
				Name:       name,
				IsFunction: true,
				Args:       args,
				Content:    value,
			}
		} else {
			defines[name] = &Define{
				Name:       name,
				IsFunction: false,
				Content:    value,
			}
		}

		definesStartIndex[name] = matches[0]

		v.Code = v.Code[:matches[0]] + v.Code[matches[len(matches)-1]:]
	}

	sortedDefines := make([]*Define, 0, len(defines))

	for _, define := range defines {
		sortedDefines = append(sortedDefines, define)
	}

	sort.Slice(sortedDefines, func(i, j int) bool {
		return definesStartIndex[sortedDefines[i].Name] > definesStartIndex[sortedDefines[j].Name]
	})

	for _, define := range sortedDefines {
		index := definesStartIndex[define.Name]

		defineRegex := regexp.MustCompile(`(?:^|[^\w])` + define.Name + `(?:[^\w]|$)`)
		v.Code = v.Code[:index] + defineRegex.ReplaceAllStringFunc(v.Code[index:], func(s string) string {
			prefix := ""
			suffix := ""

			if s[0] != define.Name[0] {
				prefix = string(s[0])
			}

			if s[len(s)-1] != define.Name[len(define.Name)-1] {
				suffix = string(s[len(s)-1])
			}

			return prefix + define.Content + suffix
		})
	}

	return v.Code, nil
}
