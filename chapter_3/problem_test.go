package chapter_3

import (
	"fmt"
	"log"
	"testing"
)

type testPoint struct {
	x int
	y int
}

// (1) *_test.go 파일명
// (2) import ("testing")
// (3) 메소드 prototype: TestXxx(t *testing.T)
func TestTwo1(t *testing.T) {
	var primeNumber uint = 223

	a := NewFiniteField(0, primeNumber)
	b := NewFiniteField(7, primeNumber)
	x1 := NewFiniteField(170, primeNumber)
	y1 := NewFiniteField(142, primeNumber)
	x2 := NewFiniteField(60, primeNumber)
	y2 := NewFiniteField(139, primeNumber)

	point1, _ := NewFiniteFieldEllipticCurve(x1, y1, a, b)
	point2, _ := NewFiniteFieldEllipticCurve(x2, y2, a, b)

	// 여기에 문제가 있음. 계산식인가?
	resultPoint, err := point1.Add(point2)

	if err != nil {
		t.Error(err)
	} else if resultPoint.x.number != 220 || resultPoint.y.number != 181 {
		t.Error("result x: 220, y:  181\n")
	}
}

func TestTwo2(t *testing.T) {
	var primeNumber uint = 223

	a := NewFiniteField(0, primeNumber)
	b := NewFiniteField(7, primeNumber)
	x1 := NewFiniteField(47, primeNumber)
	y1 := NewFiniteField(71, primeNumber)
	x2 := NewFiniteField(17, primeNumber)
	y2 := NewFiniteField(56, primeNumber)

	point1, _ := NewFiniteFieldEllipticCurve(x1, y1, a, b)
	point2, _ := NewFiniteFieldEllipticCurve(x2, y2, a, b)

	// 여기에 문제가 있음. 계산식인가?
	resultPoint, err := point1.Add(point2)

	if err != nil {
		t.Error(err)
	} else if resultPoint.x.number != 215 || resultPoint.y.number != 68 {
		t.Error("result x: 215, y:  68\n")
	}
}

func TestTwo3(t *testing.T) {
	var primeNumber uint = 223

	a := NewFiniteField(0, primeNumber)
	b := NewFiniteField(7, primeNumber)
	x1 := NewFiniteField(143, primeNumber)
	y1 := NewFiniteField(98, primeNumber)
	x2 := NewFiniteField(76, primeNumber)
	y2 := NewFiniteField(66, primeNumber)

	point1, _ := NewFiniteFieldEllipticCurve(x1, y1, a, b)
	point2, _ := NewFiniteFieldEllipticCurve(x2, y2, a, b)

	// 여기에 문제가 있음. 계산식인가?
	resultPoint, err := point1.Add(point2)

	if err != nil {
		t.Error(err)
	} else if resultPoint.x.number != 47 || resultPoint.y.number != 71 {
		t.Error("result x: 47, y: 71\n")
	}

}

func TestFour1(test *testing.T) {
	var primeNumber uint = 223

	// Constants are created at compile time
	// even when defined as locals in functions, and can only be numbers, characters (rues), strings or booleans
	answerPoints := [22]testPoint{
		{47, 71},
		{47, 71},
		{36, 111},
		{15, 137},
		{194, 51},
		{126, 96},
		{139, 137},
		{92, 47},
		{116, 55},
		{69, 86},
		{154, 150},
		{154, 73},
		{69, 137},
		{116, 168},
		{92, 176},
		{139, 86},
		{126, 127},
		{194, 172},
		{15, 86},
		{36, 112},
		{47, 152},
	}

	a := NewFiniteField(0, primeNumber)
	b := NewFiniteField(7, primeNumber)
	x := NewFiniteField(47, primeNumber)
	y := NewFiniteField(71, primeNumber)

	// 사실 이게 좋지 않은게, Pointer 는 HEAP 에 저장되서 localization 및 cache hit 가능성이 떨어짐.
	// 그러나 그대로 받아오길 하면 값전체 복사가 일어나서
	// 참조가 잦으면 Stack 하고 복사
	// 잦지 않고 한두번 쓰고 말거면 HEAP 이 괜찮음.
	point, _ := NewFiniteFieldEllipticCurve(x, y, a, b)

	var results [22]*finiteFieldEllipticCurve

	results[0] = point
	results[1] = point

	for i := 2; i <= 20; i++ {
		fmt.Printf("i : %d, ", i)

		if nextFiniteFieldEllipticCurve, err := results[i-1].Add(point); err == nil {
			results[i] = nextFiniteFieldEllipticCurve
		} else {
			// 21번째 연산에서 21 * (47, 71) == 0 이기 때문에 에러 발생
			log.Fatalln(err)
		}
	}

	for i := 1; i <= 20; i++ {
		pointX := results[i].x.number
		pointY := results[i].y.number

		if pointX != answerPoints[i].x || pointY != answerPoints[i].y {
			test.Error("Doesn't match index : ", i)
		}
	}

	fmt.Println("=======additional=======")

	//scalarResult := results[0].ScalarMultiply(20)
	//fmt.Println(*scalarResult)

}
