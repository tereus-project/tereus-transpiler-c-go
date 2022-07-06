package transpiler

import (
	"testing"

	"github.com/tereus-project/tereus-transpiler-std/core"
)

func assertTranspilation(t *testing.T, sourceCode string, targetCode string) {
	core.AssertTranspilation(t, &core.AssertTranspilationConfig{
		SourceFilename:    "test.go",
		SourceCode:        sourceCode,
		TargetCode:        targetCode,
		TranspileFunction: Transpile,
	})
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

	assertTranspilation(t, source, target)
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

	assertTranspilation(t, source, target)
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

	assertTranspilation(t, source, target)
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

	assertTranspilation(t, source, target)
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

	assertTranspilation(t, source, target)
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

	assertTranspilation(t, source, target)
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

	assertTranspilation(t, source, target)
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

	assertTranspilation(t, source, target)
}

func TestContinue(t *testing.T) {
	source := `
int main() {
	int a = 1;
	while (a < 10) {
		a++;
		if (a == 5) {
			continue;
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
			continue
		}
	}
	os.Exit(0)
}
`

	assertTranspilation(t, source, target)
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

	assertTranspilation(t, source, target)
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

	assertTranspilation(t, source, target)
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

	assertTranspilation(t, source, target)
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
	var string *int8
	string = (*int8)(libc.Malloc(int(unsafe.Sizeof((int8)(0))) * 5))
	libc.Memset((*void)(string), byte('.'), 5)
	fmt.Printf("string: %s", string)
	os.Exit(0)
}
`

	assertTranspilation(t, source, target)
}

func TestMathHLog10(t *testing.T) {
	source := `
#include <math.h>

int main() {
	int value = log10(10);
	printf("%d", value);
}
`

	target := `
package main

import (
	"fmt"
	"math"
)

func main() {
	value := int(math.Log10(float64(10)))
	fmt.Printf("%d", value)
}
`

	assertTranspilation(t, source, target)
}
