package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface{
	~int | ~int64 | ~float64
}

type IntNumber int

func genericSum[T Number] (m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func sum[T comparable] (a, b, c T) string {
	if a == b && b == c {
		return "Equal"
	}
	return "Different"
}

func max[T constraints.Ordered] (a, b T) T {
	if a > b {
		return a
	}
	return b
}

type stringer interface {
	String() string
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func concat[T stringer] (vals []T) string{
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

func main() {
	var a, b, c IntNumber
	a = 1
	b = 2
	c = 3

	fmt.Println(concat([]MyString{"con", "cat"}))

	fmt.Println(genericSum(map[string]IntNumber{"a": a, "b": b, "c": c}))
	fmt.Println(sum(a, b, c))
	fmt.Println(max(a, b))
}