package remixer

import (
	"os"
	"strings"
	"testing"
)

func testRemix(t *testing.T, source string, target string) {
	source = strings.TrimSpace(source)
	target = strings.TrimSpace(target)

	dir := t.TempDir()
	sourceFile := dir + "/test.go"
	e := os.WriteFile(sourceFile, []byte(source), 0644)
	if e != nil {
		t.Error(e)
	}

	output, e := Remix(sourceFile)
	if e != nil {
		t.Error(e)
	}

	output = strings.TrimSpace(output)
	if output != target {
		t.Errorf("Expected %s, got %s", target, output)
	}
}

func TestEmptyFunction(t *testing.T) {
	source := `
int main() {
	
}
`

	target := `
package main

func main() {

}
`

	testRemix(t, source, target)
}
