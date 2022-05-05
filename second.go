// Петя торопился в школу и неправильно написал программу, которая сначала находит
// квадраты двух чисел, а затем их суммирует. Исправьте его программу.

package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	var c int
	a = a * a
	b = b * b
	c = a + b
	fmt.Println(c)
}
