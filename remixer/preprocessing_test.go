package remixer

import (
	"os"
	"strings"
	"testing"
)

func testPreprocessor(t *testing.T, source string, target string) {
	source = strings.TrimSpace(source)
	target = strings.TrimSpace(target)

	dir := t.TempDir()
	sourceFile := dir + "/test.go"
	err := os.WriteFile(sourceFile, []byte(source), 0644)
	if err != nil {
		t.Error(err)
	}

	preprocessor, err := NewPreprocessor(sourceFile)
	if err != nil {
		t.Error(err)
	}

	output, err := preprocessor.Preprocess()
	if err != nil {
		t.Error(err)
	}

	output = strings.TrimSpace(output)
	if output != target {
		t.Errorf("expected %s, got %s", target, output)
	}
}

func TestPreprocessingSimpleSubstitution(t *testing.T) {
	source := `
#define A 1

void main() {
	int a = A;
}
`

	target := `
void main() {
	int a = 1;
}
`

	testPreprocessor(t, source, target)
}
