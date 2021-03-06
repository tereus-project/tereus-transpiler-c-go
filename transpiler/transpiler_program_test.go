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
}
`

	target := `
package main

import (
	"fmt"
	"os"

	"github.com/tereus-project/tereus-transpiler-c-go/libc"
)

func main() {
	var r int
	var c int
	var a [100][100]int
	var b [100][100]int
	var sum [100][100]int
	var i int
	var j int
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
	var i int
	isPrime := 1
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
	var c int8
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

func TestPascalTriangleProgram(t *testing.T) {
	source := `
#include <stdio.h>
int main() {
   int rows, i, j, space;
   printf("Enter the number of rows: ");
   scanf("%d", &rows);
   for (i = rows; i >= 1; --i) {
      for (space = 0; space < rows - i; ++space)
         printf("  ");
      for (j = i; j <= 2 * i - 1; ++j)
         printf("* ");
      for (j = 0; j < i - 1; ++j)
         printf("* ");
      printf("\n");
   }
   return 0;
}
`
	target := `
package main

import (
	"fmt"
	"os"

	"github.com/tereus-project/tereus-transpiler-c-go/libc"
)

func main() {
	var rows int
	var i int
	var j int
	var space int
	fmt.Printf("Enter the number of rows: ")
	fmt.Scanf("%d", &rows)
	for i = rows; i >= 1; libc.PreDec(&i) {
		for space = 0; space < rows-i; libc.PreInc(&space) {
			fmt.Printf("  ")
		}
		for j = i; j <= 2*i-1; libc.PreInc(&j) {
			fmt.Printf("* ")
		}
		for j = 0; j < i-1; libc.PreInc(&j) {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
	os.Exit(0)
}
`

	assertTranspilation(t, source, target)
}

func TestNumberPalindromeProgram(t *testing.T) {
	source := `
#include <stdio.h>
int main() {
  int n, reversed = 0, remainder, original;
    printf("Enter an integer: ");
    scanf("%d", &n);
    original = n;

    // reversed integer is stored in reversed variable
    while (n != 0) {
        remainder = n % 10;
        reversed = reversed * 10 + remainder;
        n /= 10;
    }

    // palindrome if orignal and reversed are equal
    if (original == reversed)
        printf("%d is a palindrome.", original);
    else
        printf("%d is not a palindrome.", original);

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
	var n int
	reversed := 0
	var remainder int
	var original int
	fmt.Printf("Enter an integer: ")
	fmt.Scanf("%d", &n)
	original = n
	// reversed integer is stored in reversed variable
	for n != 0 {
		remainder = n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}
	// palindrome if orignal and reversed are equal
	if original == reversed {
		fmt.Printf("%d is a palindrome.", original)
	} else {
		fmt.Printf("%d is not a palindrome.", original)
	}
	os.Exit(0)
}
`

	assertTranspilation(t, source, target)
}

func TestTransposeMatrixProgram(t *testing.T) {
	source := `
#include <stdio.h>
int main() {
	int a[10][10], transpose[10][10], r, c;
	printf("Enter rows and columns: ");
	scanf("%d %d", &r, &c);

	// asssigning elements to the matrix
	printf("\nEnter matrix elements:\n");
	for (int i = 0; i < r; ++i)
	for (int j = 0; j < c; ++j) {
	printf("Enter element a%d%d: ", i + 1, j + 1);
	scanf("%d", &a[i][j]);
	}

	// printing the matrix a[][]
	printf("\nEntered matrix: \n");
	for (int i = 0; i < r; ++i)
	for (int j = 0; j < c; ++j) {
	printf("%d  ", a[i][j]);
	if (j == c - 1)
	printf("\n");
	}

	// computing the transpose
	for (int i = 0; i < r; ++i)
	for (int j = 0; j < c; ++j) {
	transpose[j][i] = a[i][j];
	}

	// printing the transpose
	printf("\nTranspose of the matrix:\n");
	for (int i = 0; i < c; ++i)
	for (int j = 0; j < r; ++j) {
	printf("%d  ", transpose[i][j]);
	if (j == r - 1)
	printf("\n");
	}
	return 0;
}
`

	target := `
package main

import (
	"fmt"
	"os"

	"github.com/tereus-project/tereus-transpiler-c-go/libc"
)

func main() {
	var a [10][10]int
	var transpose [10][10]int
	var r int
	var c int
	fmt.Printf("Enter rows and columns: ")
	fmt.Scanf("%d %d", &r, &c)
	// asssigning elements to the matrix
	fmt.Printf("\nEnter matrix elements:\n")
	for i := 0; i < r; libc.PreInc(&i) {
		for j := 0; j < c; libc.PreInc(&j) {
			fmt.Printf("Enter element a%d%d: ", i+1, j+1)
			fmt.Scanf("%d", &a[i][j])
		}
	}
	// printing the matrix a[][]
	fmt.Printf("\nEntered matrix: \n")
	for i := 0; i < r; libc.PreInc(&i) {
		for j := 0; j < c; libc.PreInc(&j) {
			fmt.Printf("%d  ", a[i][j])
			if j == c-1 {
				fmt.Printf("\n")
			}
		}
	}
	// computing the transpose
	for i := 0; i < r; libc.PreInc(&i) {
		for j := 0; j < c; libc.PreInc(&j) {
			transpose[j][i] = a[i][j]
		}
	}
	// printing the transpose
	fmt.Printf("\nTranspose of the matrix:\n")
	for i := 0; i < c; libc.PreInc(&i) {
		for j := 0; j < r; libc.PreInc(&j) {
			fmt.Printf("%d  ", transpose[i][j])
			if j == r-1 {
				fmt.Printf("\n")
			}
		}
	}
	os.Exit(0)
}
`

	assertTranspilation(t, source, target)
}

func TestPrimeNumberBetweenTwoNumbersProgram(t *testing.T) {
	source := `
#include <stdio.h>

// user-defined function to check prime number
int checkPrimeNumber(int n) {
	int j, flag = 1;

	for (j = 2; j <= n / 2; ++j) {

	if (n % j == 0) {
		flag = 0;
		break;
	}
	}

	return flag;
}

int main() {

	int n1, n2, i, flag;

	printf("Enter two positive integers: ");
	scanf("%d %d", &n1, &n2);

	// swap n1 and n2 if n1 > n2
	if (n1 > n2) {
	n1 = n1 + n2;
	n2 = n1 - n2;
	n1 = n1 - n2;
	}

	printf("Prime numbers between %d and %d are: ", n1, n2);
	for (i = n1 + 1; i < n2; ++i) {

	// flag will be equal to 1 if i is prime
	flag = checkPrimeNumber(i);

	if (flag == 1) {
		printf("%d ", i);
	}
	}

	return 0;
}
`
	target := `
package main

import (
	"fmt"
	"os"

	"github.com/tereus-project/tereus-transpiler-c-go/libc"
)

// user-defined function to check prime number

func checkPrimeNumber(n int) int {
	var j int
	flag := 1
	for j = 2; j <= n/2; libc.PreInc(&j) {
		if n%j == 0 {
			flag = 0
			break
		}
	}
	return flag
}

func main() {
	var n1 int
	var n2 int
	var i int
	var flag int
	fmt.Printf("Enter two positive integers: ")
	fmt.Scanf("%d %d", &n1, &n2)
	// swap n1 and n2 if n1 > n2
	if n1 > n2 {
		n1 = n1 + n2
		n2 = n1 - n2
		n1 = n1 - n2
	}
	fmt.Printf("Prime numbers between %d and %d are: ", n1, n2)
	for i = n1 + 1; i < n2; libc.PreInc(&i) {
		// flag will be equal to 1 if i is prime
		flag = checkPrimeNumber(i)
		if flag == 1 {
			fmt.Printf("%d ", i)
		}
	}
	os.Exit(0)
}
`

	assertTranspilation(t, source, target)
}

func TestHashmapProgram(t *testing.T) {
	source := `
#include <stdio.h>
#include <stdlib.h>
struct node
{
    int key;
    int val;
    struct node *next;
};
struct table
{
    int size;
    struct node **list;
};
struct table *createTable(int size)
{
    struct table *t = (struct table *)malloc(sizeof(struct table));
    t->size = size;
    t->list = (struct node **)malloc(sizeof(struct node *) * size);
    int i;
    for (i = 0; i < size; i++)
        t->list[i] = NULL;
    return t;
}
int hashCode(struct table *t, int key)
{
    if (key < 0)
        return -(key % t->size);
    return key % t->size;
}
void insert(struct table *t, int key, int val)
{
    int pos = hashCode(t, key);
    struct node *list = t->list[pos];
    struct node *newNode = (struct node *)malloc(sizeof(struct node));
    struct node *temp = list;
    while (temp)
    {
        if (temp->key == key)
        {
            temp->val = val;
            return;
        }
        temp = temp->next;
    }
    newNode->key = key;
    newNode->val = val;
    newNode->next = list;
    t->list[pos] = newNode;
}
int lookup(struct table *t, int key)
{
    int pos = hashCode(t, key);
    struct node *list = t->list[pos];
    struct node *temp = list;
    while (temp)
    {
        if (temp->key == key)
        {
            return temp->val;
        }
        temp = temp->next;
    }
    return -1;
}
int main()
{
    struct table *t = createTable(5);
    insert(t, 2, 3);
    printf("%d", lookup(t, 2));
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

type node struct {
	Key  int
	Val  int
	Next *node
}

type table struct {
	Size int
	List **node
}

func createTable(size int) *table {
	t := (*table)(libc.Malloc(int(unsafe.Sizeof(table{}))))
	t.Size = size
	t.List = (**node)(libc.Malloc(int(unsafe.Sizeof((*node)(nil))) * size))
	var i int
	for i = 0; i < size; i++ {
		*(**node)(unsafe.Add(unsafe.Pointer(t.List), uintptr(i)*unsafe.Sizeof((**node)(nil)))) = (*node)(nil)
	}
	return t
}

func hashCode(t *table, key int) int {
	if key < 0 {
		return -(key % t.Size)
	}
	return key % t.Size
}

func insert(t *table, key int, val int) {
	pos := hashCode(t, key)
	list := *(**node)(unsafe.Add(unsafe.Pointer(t.List), uintptr(pos)*unsafe.Sizeof((**node)(nil))))
	newNode := (*node)(libc.Malloc(int(unsafe.Sizeof(node{}))))
	temp := list
	for temp != nil {
		if temp.Key == key {
			temp.Val = val
			return
		}
		temp = temp.Next
	}
	newNode.Key = key
	newNode.Val = val
	newNode.Next = list
	*(**node)(unsafe.Add(unsafe.Pointer(t.List), uintptr(pos)*unsafe.Sizeof((**node)(nil)))) = newNode
}

func lookup(t *table, key int) int {
	pos := hashCode(t, key)
	list := *(**node)(unsafe.Add(unsafe.Pointer(t.List), uintptr(pos)*unsafe.Sizeof((**node)(nil))))
	temp := list
	for temp != nil {
		if temp.Key == key {
			return temp.Val
		}
		temp = temp.Next
	}
	return -1
}

func main() {
	t := createTable(5)
	insert(t, 2, 3)
	fmt.Printf("%d", lookup(t, 2))
	os.Exit(0)
}
`

	assertTranspilation(t, source, target)
}
