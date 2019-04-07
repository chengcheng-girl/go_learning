package normal

import (
	"fmt"
	"testing"
)
func init() {
	fmt.Println("init2")
}

func init() {
	fmt.Println("init1")
}


func Square(n int) int {
	return n * n
}

func TestSquare(t *testing.T)  {
	fmt.Print(Square(10))
}


func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
