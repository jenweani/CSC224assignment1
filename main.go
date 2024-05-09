package main
import (
	"fmt"
	"math"
)
func NRT(n int) int {
	if n < 0 {
		panic("n cannot be negative")
	}

	t := make([]int, n+1)

	t[0] = 1
	t[1] = 3

	for i := 2; i <= n; i++ {
		t[i] = 4*t[i-1] - 3*t[i-2] + 4
	}

	return t[n]
}

func RT(n int) int {
	if (n == 1) {
		return 3;
	}
	if (n == 0) {
		return 1;
	}else{
		return (4 * RT(n - 1)) - (3 * RT(n - 2)) + 4;
	}
}

func T(n int) int {
	return (2 * int(math.Pow(3, float64(n))) - (2 * n) - 1);
}

func main() {
	n := 5;
	Tres := T(n);
	NRTres := NRT(n);
	RTres := RT(n);
	fmt.Println("T result:", Tres);
	fmt.Println("NRT result:", NRTres);
	fmt.Println("RT result:", RTres);
}