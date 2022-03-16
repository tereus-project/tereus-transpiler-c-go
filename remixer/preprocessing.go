package remixer

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/dlclark/regexp2"
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
			return "", fmt.Errorf("macro '%s' already defined", name)
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

		if define.IsFunction {
			defineRegex := regexp.MustCompile(`(?:^|[^\w])` + define.Name + `\s*\(`)

			for matches := defineRegex.FindAllStringSubmatchIndex(v.Code, -1); matches != nil; matches = defineRegex.FindAllStringSubmatchIndex(v.Code, -1) {
				match := matches[len(matches)-1]

				args := []string{""}

				start := match[0]
				end := match[1]

				for parenStack, isInString := 1, false; end < len(v.Code); end++ {
					c := v.Code[end]

					if !isInString {
						if c == '(' {
							parenStack++
						} else if c == ')' {
							parenStack--

							if parenStack == 0 {
								break
							}
						} else if c == '\'' || c == '"' {
							isInString = true
						} else if c == ',' && parenStack == 1 {
							args = append(args, "")
							continue
						}
					} else if c == '\\' && len(v.Code) > end+1 {
						args[len(args)-1] += string(c) + string(v.Code[end+1])
						end++
						continue
					} else if c == '"' || c == '\'' {
						isInString = false
					}

					args[len(args)-1] += string(c)
				}

				if len(args) == 1 && args[0] == "" {
					args = args[1:]
				}

				if len(args) != len(define.Args) {
					return "", fmt.Errorf("invalid number of arguments for macro '%s': %d of %d expected", define.Name, len(args), len(define.Args))
				}

				argsPosition := make(map[string]int)

				for i, arg := range define.Args {
					argsPosition[arg] = i
				}

				defineArgumentRegex := regexp2.MustCompile(`(?<=^|[^\w])(`+strings.Join(define.Args, "|")+`)(?=[^\w]|$)`, 0)

				content, err := defineArgumentRegex.ReplaceFunc(define.Content, func(match regexp2.Match) string {
					return strings.TrimSpace(args[argsPosition[match.String()]])
				}, -1, -1)

				if err != nil {
					return "", err
				}

				prefix := ""

				if v.Code[start] != define.Name[0] {
					prefix = string(v.Code[start])
				}

				v.Code = prefix + v.Code[:start+1] + content + v.Code[end+1:]
			}
		} else {
			defineRegex := regexp2.MustCompile(`(?<=^|[^\w])`+define.Name+`(?=[^\w]|$)`, 0)

			content, err := defineRegex.Replace(v.Code[index:], define.Content, -1, -1)
			if err != nil {
				return "", err
			}

			v.Code = v.Code[:index] + content
		}
	}

	return v.Code, nil
}
