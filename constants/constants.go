package main;

import (
	"math"
	"fmt"
)

const s string = "constant";

func main() {
	fmt.Println(s);

	const n = 500000000;

	const d = 2e20 / n;

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n));
}