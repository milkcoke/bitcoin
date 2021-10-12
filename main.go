package main

import (
	"fmt"
	"github.com/milkcoke/bitcoin/chapter_1"
)

// https://jacking75.github.io/go_heap-allocations/
// Go 컴파일러는 코드를 분석하여 가능하면 스택할당, 그렇지 않으면 힙 할당된다.
// 스택 할당은 변수의 수명과 메모리 사용량이 컴파일 시 확정이 가능한 경우.
// 그 외의 경우엔 Heap 할당됨.
// Stack < Heap Allocation 더 무거움. (stack allocation 은 CPU 명령어 2개로 끝남, 할당 1개 해제 1개)

func main() {
	// https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go
	// golang 의 modulo 연산은 음수를 뱉는다.
	//fmt.Println((-5) % 3)

	// order: 위수
	var order uint = 57
	var order2 uint = 13
	var order3 uint = 31

	a := chapter_1.NewFiniteField(44, order)
	b := chapter_1.NewFiniteField(33, order)
	c := chapter_1.NewFiniteField(9, order)
	d := chapter_1.NewFiniteField(29, order)
	e := chapter_1.NewFiniteField(3, order2)
	f := chapter_1.NewFiniteField(12, order2)
	g := chapter_1.NewFiniteField(10, order2)

	dividend := chapter_1.NewFiniteField(3, order3)
	divisor := chapter_1.NewFiniteField(24, order3)
	result1, _ := chapter_1.Add(a, b)
	result2, _ := chapter_1.Subtract(c, d)
	result3, _ := chapter_1.Multiply(e, f)
	result4, _ := chapter_1.Divide(dividend, divisor)
	fmt.Printf("result1 : %d, result2: %d\n", result1.GetNumber(), result2.GetNumber())
	// %t is bool format in go
	fmt.Printf("result3: %t\n", result3.GetNumber() == g.GetNumber())

	fmt.Printf("result4: %d", result4.GetNumber())
}
