package auto_documentation

import "fmt"

func ExampleSum() {
	x := Sum(1, 2)
	fmt.Println(x)
	fmt.Println(Sum(-1, -1))
	fmt.Println(Sum(0, 0))

	// Output:
	// 3
	// -2
	// 0
}
