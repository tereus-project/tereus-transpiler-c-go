package transpiler

import (
	"testing"
)

func TestMatrixProgram(t *testing.T) {
	source := `
#include <stdio.h>
int main() {
  int r, c, a[100][100], b[100][100], sum[100][100], i, j;
  printf("Enter the number of rows (between 1 and 100): ");
  scanf("%d", &r);
  printf("Enter the number of columns (between 1 and 100): ");
  scanf("%d", &c);

  printf("\nEnter elements of 1st matrix:\n");
  for (i = 0; i < r; ++i)
    for (j = 0; j < c; ++j) {
      printf("Enter element a%d%d: ", i + 1, j + 1);
      scanf("%d", &a[i][j]);
    }

  printf("Enter elements of 2nd matrix:\n");
  for (i = 0; i < r; ++i)
    for (j = 0; j < c; ++j) {
      printf("Enter element b%d%d: ", i + 1, j + 1);
      scanf("%d", &b[i][j]);
    }

  // adding two matrices
  for (i = 0; i < r; ++i)
    for (j = 0; j < c; ++j) {
      sum[i][j] = a[i][j] + b[i][j];
    }

  // printing the result
  printf("\nSum of two matrices: \n");
  for (i = 0; i < r; ++i)
    for (j = 0; j < c; ++j) {
      printf("%d   ", sum[i][j]);
      if (j == c - 1) {
        printf("\n\n");
      }
    }

  return 0;
}`

	target := `
package main

import (
	"fmt"
	"os"

	"github.com/tereus-project/tereus-transpiler-c-go/libc"
)

func main() {
	r, c, a, b, sum, i, j := 0, 0, [100][100]int{}, [100][100]int{}, [100][100]int{}, 0, 0
	fmt.Printf("Enter the number of rows (between 1 and 100): ")
	fmt.Scanf("%d", &r)
	fmt.Printf("Enter the number of columns (between 1 and 100): ")
	fmt.Scanf("%d", &c)
	fmt.Printf("\nEnter elements of 1st matrix:\n")
	for i = 0; i < r; libc.PreInc(&i) {
		for j = 0; j < c; libc.PreInc(&j) {
			fmt.Printf("Enter element a%d%d: ", i+1, j+1)
			fmt.Scanf("%d", &a[i][j])
		}
	}
	fmt.Printf("Enter elements of 2nd matrix:\n")
	for i = 0; i < r; libc.PreInc(&i) {
		for j = 0; j < c; libc.PreInc(&j) {
			fmt.Printf("Enter element b%d%d: ", i+1, j+1)
			fmt.Scanf("%d", &b[i][j])
		}
	}
	// adding two matrices
	for i = 0; i < r; libc.PreInc(&i) {
		for j = 0; j < c; libc.PreInc(&j) {
			sum[i][j] = a[i][j] + b[i][j]
		}
	}
	// printing the result
	fmt.Printf("\nSum of two matrices: \n")
	for i = 0; i < r; libc.PreInc(&i) {
		for j = 0; j < c; libc.PreInc(&j) {
			fmt.Printf("%d   ", sum[i][j])
			if j == c-1 {
				fmt.Printf("\n\n")
			}
		}
	}
	os.Exit(0)
}
`

	assertTranspilation(t, source, target)
}

func TestPrimeProgram(t *testing.T) {
	source := `
#include <stdio.h>

// function to check prime number
int checkPrime(int n) {
	int i, isPrime = 1;

	// 0 and 1 are not prime numbers
	if (n == 0 || n == 1) {
	isPrime = 0;
	}
	else {
	for(i = 2; i <= n/2; ++i) {
		if(n % i == 0) {
		isPrime = 0;
		break;
		}
	}
	}

	return isPrime;
}

int main() {
	printf("7 is prime: %d\n", checkPrime(7));
	printf("8 is prime: %d\n", checkPrime(8));
}
`

	target := `
package main

import (
	"fmt"

	"github.com/tereus-project/tereus-transpiler-c-go/libc"
)

// function to check prime number

func checkPrime(n int) int {
	i, isPrime := 0, 1
	// 0 and 1 are not prime numbers
	if n == 0 || n == 1 {
		isPrime = 0
	} else {
		for i = 2; i <= n/2; libc.PreInc(&i) {
			if n%i == 0 {
				isPrime = 0
				break
			}
		}
	}
	return isPrime
}

func main() {
	fmt.Printf("7 is prime: %d\n", checkPrime(7))
	fmt.Printf("8 is prime: %d\n", checkPrime(8))
}
`

	assertTranspilation(t, source, target)
}

func TestReverseStringProgram(t *testing.T) {
	source := `
#include <stdio.h>

void reverseSentence() {
	char c;
	scanf("%c", &c);
	if (c != '\n') {
		reverseSentence();
		printf("%c", c);
	}
}

int main() {
	printf("Enter a sentence: ");
	reverseSentence();
	return 0;
}
`

	target := `
package main

import (
	"fmt"
	"os"
)

func reverseSentence() {
	c := byte(0)
	fmt.Scanf("%c", &c)
	if c != '\n' {
		reverseSentence()
		fmt.Printf("%c", c)
	}
}

func main() {
	fmt.Printf("Enter a sentence: ")
	reverseSentence()
	os.Exit(0)
}
`

	assertTranspilation(t, source, target)
}
