package main

import (
	"fmt"
	"github.com/milkcoke/bitcoin/chapter_2"
	"log"
)

// https://jacking75.github.io/go_heap-allocations/
// Go 컴파일러는 코드를 분석하여 가능하면 스택할당, 그렇지 않으면 힙 할당된다.
// 스택 할당은 변수의 수명과 메모리 사용량이 컴파일 시 확정이 가능한 경우.
// 그 외의 경우엔 Heap 할당됨.
// Stack < Heap Allocation 더 무거움. (stack allocation 은 CPU 명령어 2개로 끝남, 할당 1개 해제 1개)

func main() {

	pointA, err1 := chapter_2.NewEllipticCurve(-1, -1, 5, 7)
	pointB, err2 := chapter_2.NewEllipticCurve(-1, -1, 5, 7)
	if err1 != nil || err2 != nil {
		log.Fatalln(err1, err2)
	}
	resultPoint, _ := chapter_2.Add(pointA, pointB)
	fmt.Println(*resultPoint)

}
