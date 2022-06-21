package transpiler

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
	preprocessorRegex         = regexp2.MustCompile(`(?<String>\"(?:(?:\\.)|[^\"])*\")|(^[ \t]*#\s*define\s+(?<Name>\w+)(?<Args>\(\s*\w+(?:\s*,\s*\w+)*\s*\))?\s+(?<Value>(?:(?:[\\][\r\n])|.)+)$)`, regexp2.Multiline)
	preprocessorArgumentRegex = regexp.MustCompile(`^[a-zA-Z_]\w*$`)
)

func (v *Preprocessor) Preprocess() (string, error) {
	defines := make(map[string]*Define)
	definesStartIndex := make(map[string]int)

	for match, _ := preprocessorRegex.FindStringMatch(v.Code); match != nil; match, _ = preprocessorRegex.FindStringMatch(v.Code) {
		if str := match.GroupByName("String"); str != nil && str.Length > 0 {
			continue
		}

		name := match.GroupByName("Name").String()
		value := match.GroupByName("Value").String()
		argsString := match.GroupByName("Args").String()

		if _, ok := defines[name]; ok {
			return "", fmt.Errorf("macro '%s' already defined", name)
		}

		if len(argsString) > 0 {
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

		definesStartIndex[name] = match.Index

		v.Code = v.Code[:match.Index] + v.Code[match.Index+match.Length:]
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
			defineRegex := regexp2.MustCompile(`(?:\"(?:(?:\\.)|[^\"])*\")|(?<Target>(?<=^|[^\w])`+define.Name+`\s*\()`, regexp2.RightToLeft)

			selectedPart := v.Code[index:]

			for match, _ := defineRegex.FindStringMatch(selectedPart); match != nil; match, _ = defineRegex.FindNextMatch(match) {
				group := match.GroupByName("Target")
				if group == nil || group.Length == 0 {
					continue
				}

				start := group.Capture.Index
				end := start + group.Length

				args := []string{""}

				for parenStack, isInString := 1, false; end < len(selectedPart); end++ {
					c := selectedPart[end]

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
					} else if c == '\\' && len(selectedPart) > end+1 {
						args[len(args)-1] += string(c) + string(selectedPart[end+1])
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

				defineArgumentRegex := regexp2.MustCompile(`(?:\"(?:(?:\\.)|[^\"])*\")|(?:(?<=^|[^\w])(?<Target>(?:`+strings.Join(define.Args, ")|(")+`))(?=[^\w]|$))`, 0)

				content, err := defineArgumentRegex.ReplaceFunc(define.Content, func(match regexp2.Match) string {
					if group := match.GroupByName("Target"); group != nil && group.Length > 0 {
						if index, ok := argsPosition[group.String()]; ok {
							return strings.TrimSpace(args[index])
						}
					}

					return match.String()
				}, -1, -1)

				if err != nil {
					return "", err
				}

				selectedPart = selectedPart[:start] + content + selectedPart[end+1:]
			}

			v.Code = v.Code[:index] + selectedPart
		} else {
			defineRegex := regexp2.MustCompile(`(?:\"(?:(?:\\.)|[^\"])*\")|(?<Target>(?<=^|[^\w])`+define.Name+`(?=[^\w]|$))`, 0)

			content, err := defineRegex.ReplaceFunc(v.Code[index:], func(match regexp2.Match) string {
				if group := match.GroupByName("Target"); group != nil && group.Length > 0 {
					return define.Content
				}

				return match.String()
			}, -1, -1)
			if err != nil {
				return "", err
			}

			v.Code = v.Code[:index] + content
		}
	}

	return v.Code, nil
}
