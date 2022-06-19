package transpiler

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
	err := os.WriteFile(sourceFile, []byte(source), 0644)
	if err != nil {
		t.Error(err)
	}

	output, err := Remix(sourceFile)
	if err != nil {
		t.Error(err)
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

func TestMainWithArgcAndArgv(t *testing.T) {
	source := `
int main(int argc, char **argv) {

}
`

	target := `
package main

import (
	"os"
)

func main() {
	argc := len(os.Args)
	argv := os.Args
}
`

	testRemix(t, source, target)
}

func TestFunctionCall(t *testing.T) {
	source := `
int add(int a, int b) {
	return a + b;
}

int main() {
	add(1, 2);
}
`

	target := `
package main

func add(a int, b int) int {
	return a + b
}

func main() {
	add(1, 2)
}
`

	testRemix(t, source, target)
}

func TestIfCondition(t *testing.T) {
	source := `
int main() {
	int a = 1;

	if (a >= 1) {
		a = 2;
	}

	return 0;
}
`

	target := `
package main

import (
	"os"
)

func main() {
	a := 1
	if a >= 1 {
		a = 2
	}
	os.Exit(0)
}
`

	testRemix(t, source, target)
}

func TestUnaryPreExpression(t *testing.T) {
	source := `
int main() {
	int a = -1;
	return 0;
}
`

	target := `
package main

import (
	"os"
)

func main() {
	a := -1
	os.Exit(0)
}
`

	testRemix(t, source, target)
}

func TestUnaryPostExpression(t *testing.T) {
	source := `
int main() {
	int a = 1;
	a++;
	return 0;
}
`

	target := `
package main

import (
	"os"
)

func main() {
	a := 1
	a++
	os.Exit(0)
}
`

	testRemix(t, source, target)
}

func TestFmtPrintf(t *testing.T) {
	source := `
int main() {
	int a = 1;
	printf("a: %d", a);
	return 0;
}
`

	target := `
package main

import (
	"fmt"
	"os"
)

func main() {
	a := 1
	fmt.Printf("a: %d", a)
	os.Exit(0)
}
`

	testRemix(t, source, target)
}

func TestBreak(t *testing.T) {
	source := `
int main() {
	int a = 1;
	while (a < 10) {
		a++;
		if (a == 5) {
			break;
		}
	}
	return 0;
}
`

	target := `
package main

import (
	"os"
)

func main() {
	a := 1
	for a < 10 {
		a++
		if a == 5 {
			break
		}
	}
	os.Exit(0)
}
`

	testRemix(t, source, target)
}

func TestStdMalloc(t *testing.T) {
	source := `
#include <stdlib.h>

int main()
{
	int *a = (int *)malloc(sizeof(int) * 5);

	return 0;
}
`

	target := `
package main

import (
	"os"
	"unsafe"

	"github.com/tereus-project/tereus-transpiler-c-go/libc"
)

func main() {
	a := (*int)(libc.Malloc(int(unsafe.Sizeof(int(0))) * 5))
	os.Exit(0)
}
`

	testRemix(t, source, target)
}

func TestGoto(t *testing.T) {
	source := `
#include <stdio.h>
int main()
{
	goto out;
out:
	printf("hey goto");
}
`

	target := `
package main

import (
	"fmt"
)

func main() {
	goto out
out:
	fmt.Printf("hey goto")
}
`

	testRemix(t, source, target)
}

func TestAssert(t *testing.T) {
	source := `
#include <stdio.h>
#include <assert.h>

void main() {
	assert(1 == 1);
}
`

	target := `
package main

import (
	"github.com/tereus-project/tereus-transpiler-c-go/libc"
)

func main() {
	libc.Assert(1 == 1)
}
`

	testRemix(t, source, target)
}

func TestMemset(t *testing.T) {
	source := `
#include <stdio.h>
#include <stdlib.h>

int main()
{
	char *string;
	string = malloc(sizeof(char) * 5);

	memset(string, '.', 5);

	printf("string: %s", string);

	return 0;
}
`

	target := `
package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/tereus-project/tereus-transpiler-c-go/libc"
)

func main() {
	string := (*int8)(nil)
	unsafe.Add(unsafe.Pointer(string), (*int8)(libc.Malloc(int(unsafe.Sizeof((int8)(0)))*5)))
	libc.Memset((*void)(string), byte('.'), 5)
	fmt.Printf("string: %s", string)
	os.Exit(0)
}
`

	testRemix(t, source, target)
}
