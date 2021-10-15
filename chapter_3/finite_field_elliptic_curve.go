package chapter_3

import (
	"errors"
	"fmt"
	"math"
)

var errInvalidParameter error = errors.New("invalid parameter error")
var errDifferentPrime error = errors.New("two finite field has different prime number each other")
var errInfiniteValue error = errors.New("invalid parameter (x or y is infinite value)")
var errNotOnSameGraph error = errors.New("not on same graph")

// chapter 1 유한체와
// chapter 2 타원곡선을 함께!
// chapter 3 은 유한체 타원곡선 암호

type finiteField struct {
	number int
	prime  uint
}

type finiteFieldEllipticCurve struct {
	x finiteField
	y finiteField
	a finiteField
	b finiteField
}

func NewFiniteField(number int, primeNumber uint) *finiteField {
	// 이거 return 한거 자동으로 Heap 에 할당되는 듯 `New` 키워드 쓴 것 처럼 (js, python, java 처럼)
	// 그렇지 않으면 main 에서 instance 받아서 계속 쓰는게 말이 안됨.
	return &finiteField{number, primeNumber}
}

// NewFiniteFieldEllipticCurve is constructor
func NewFiniteFieldEllipticCurve(x, y, a, b *finiteField) (*finiteFieldEllipticCurve, error) {
	positiveInfiniteValue := math.MaxInt
	negativeInfiniteValue := math.MinInt

	if x.number == positiveInfiniteValue || x.number == negativeInfiniteValue || y.number == positiveInfiniteValue || y.number == negativeInfiniteValue {
		return nil, errInfiniteValue
	}
	// y^2 = x^3 + ax + b
	// Add((Pow(&x, 3), Multiply(&a, &x))), b
	powered, _ := Pow(x, 3)
	multiplied, _ := Multiply(a, x)
	added, _ := Add(powered, multiplied)
	xResult, _ := Add(added, b)
	yResult, _ := Pow(y, 2)

	if *yResult != *xResult {
		fmt.Printf("Failed to create ellipticCurve on coordinate (x:%d, y:%d)\n", x.number, y.number)
		return nil, errInvalidParameter
	} else {
		fmt.Printf("succeed to create ellipticCurve on coordinate (x:%d, y:%d)\n", x.number, y.number)
		return &finiteFieldEllipticCurve{*x, *y, *a, *b}, nil
	}
}

func Add(aField, bField *finiteField) (*finiteField, error) {

	if aField.prime != bField.prime {
		return nil, errDifferentPrime
	}

	// in Go, modulo return negative value when dividend is negative number
	// so add prime number for casting to positive remain number
	resultNumber := uint(aField.number+bField.number+int(aField.prime)) % aField.prime

	// almost cases in go, pointer is allocated in HEAP segment
	// For memory localization (cache hit rate), some developers recommend using stack allocation (don't use pointer)
	resultFiniteField := &finiteField{int(resultNumber), aField.prime}

	return resultFiniteField, nil
}

func Subtract(aField, bField *finiteField) (*finiteField, error) {
	if aField.prime != bField.prime {
		return nil, errDifferentPrime
	}

	// plus prime until positive value
	var minusedValue int = aField.number - bField.number
	var resultNumber int = 0

	if minusedValue < 0 {
		absValue := int(math.Abs(float64(minusedValue)))
		// 절대값 몫을 구함
		var quotient int = absValue / int(aField.prime)
		// 절대값 만큼 곱해 더하고 한번 더 더하면 양수로 전환됨.
		minusedValue += quotient*int(aField.prime) + int(aField.prime)
		// 그 값에 modulo 연산하면 됨.
		resultNumber = minusedValue % int(aField.prime)
	} else {
		// 양수 케이스는 바로 modulo 연산.
		resultNumber = int(uint(minusedValue) % aField.prime)
	}

	resultFiniteField := &finiteField{resultNumber, aField.prime}

	return resultFiniteField, nil
}

func Multiply(aField, bField *finiteField) (*finiteField, error) {

	if aField.prime != bField.prime {
		return nil, errDifferentPrime
	}

	var multipliedNumber int = aField.number * bField.number
	var resultNumber int = 0

	// processing negative value
	if resultNumber < 0 {
		abs := int(math.Abs(float64(multipliedNumber)))
		var quotient int = abs / int(aField.prime)
		multipliedNumber += quotient*int(aField.prime) + int(aField.prime)
		resultNumber %= int(aField.prime)
	} else {
		resultNumber = int(uint(multipliedNumber) % aField.prime)
	}

	return &finiteField{resultNumber, aField.prime}, nil

}

// Divide
func Divide(aField, bField *finiteField) (*finiteField, error) {
	poweredNumber, _ := Pow(bField, int(aField.prime-2))
	var resultNumber int = aField.number * poweredNumber.number % int(aField.prime)

	return &finiteField{resultNumber, aField.prime}, nil
}

// Pow is just for n^exp % exp
func Pow(aField *finiteField, exponent int) (*finiteField, error) {
	// change exponent to value in [0, p-2]
	// Formular for modulo deeper into the equation to keep the values low at every step
	// (a + b) % n = (a % n + b % n) % n
	// (a * b) % n = (a % n * b * n) % n
	// https://stackoverflow.com/questions/52283446/raising-a-number-to-a-huge-exponent
	// https://www.geeksforgeeks.org/modular-exponentiation-power-in-modular-arithmetic/
	//resultNumber := int(math.Pow(float64(aField.number), float64(exponent))) % int(aField.prime)

	// Optimization

	// initialize answer
	result := 1

	// update number if it's more than or equal to exponent
	number := aField.number % int(aField.prime)

	for exponent > 0 {
		if exponent&1 == 1 {
			result = result * number % int(aField.prime)
		}
		// exp = exp / 2
		exponent = exponent >> 1

		// Change number to number ^ 2
		number = number * number % int(aField.prime)
	}

	return NewFiniteField(result, aField.prime), nil
}

func (thisFiniteFieldEllipticCurve *finiteFieldEllipticCurve) Add(bFiniteFieldEllipticCurve *finiteFieldEllipticCurve) (*finiteFieldEllipticCurve, error) {
	// 그래프가 다른경우
	if thisFiniteFieldEllipticCurve.a != bFiniteFieldEllipticCurve.a || thisFiniteFieldEllipticCurve.b != bFiniteFieldEllipticCurve.b {
		return nil, errNotOnSameGraph
		// 1. x축에 수직인 직선 (무한원점)
		// 그냥 무한원점 (x = a a는 상수) 를 x 무한대로 표현함.
	} else if thisFiniteFieldEllipticCurve.x.number == math.MaxInt || thisFiniteFieldEllipticCurve.x.number == math.MinInt {
		return bFiniteFieldEllipticCurve, nil
	} else if bFiniteFieldEllipticCurve.x.number == math.MaxInt || bFiniteFieldEllipticCurve.x.number == math.MinInt {
		return thisFiniteFieldEllipticCurve, nil

		// 1-1. 역원을 더하는 케이스, x 좌표는 서로 같다.
	} else if thisFiniteFieldEllipticCurve.x == bFiniteFieldEllipticCurve.x && thisFiniteFieldEllipticCurve.y != bFiniteFieldEllipticCurve.y {

		infiniteOriginPointX := NewFiniteField(math.MaxInt, thisFiniteFieldEllipticCurve.x.prime)
		infiniteOriginPointY := NewFiniteField(math.MaxInt, thisFiniteFieldEllipticCurve.y.prime)

		// 무한 원점 반환
		return NewFiniteFieldEllipticCurve(
			infiniteOriginPointX,
			infiniteOriginPointY,
			&thisFiniteFieldEllipticCurve.a,
			&thisFiniteFieldEllipticCurve.b,
		)

		// 3. P1 == P2 동일한 점에 대한 덧셈 구하기 (접선의 교점 && x축 대칭)
	} else if thisFiniteFieldEllipticCurve.x == bFiniteFieldEllipticCurve.x && thisFiniteFieldEllipticCurve.y == bFiniteFieldEllipticCurve.y {

		if thisFiniteFieldEllipticCurve.y.number != 0 {

			// s is inclination of tangent line
			// (3x^2 + a) / 2y
			// 3x
			point3 := NewFiniteField(3, thisFiniteFieldEllipticCurve.x.prime)
			point2 := NewFiniteField(2, thisFiniteFieldEllipticCurve.x.prime)
			powed, _ := Pow(&thisFiniteFieldEllipticCurve.x, 2)
			multiplied, _ := Multiply(powed, point3)
			added, _ := Add(multiplied, &thisFiniteFieldEllipticCurve.a)
			multiplied2, _ := Multiply(point2, &thisFiniteFieldEllipticCurve.y)
			resultInclination, _ := Divide(added, multiplied2)

			var s *finiteField = resultInclination
			// 여기 뭔가 틀린듯? 10 - 18
			powed, _ = Pow(s, 2)
			multiplied, _ = Multiply(&thisFiniteFieldEllipticCurve.x, point2)
			pointX, _ := Subtract(powed, multiplied)

			subtracted, _ := Subtract(&thisFiniteFieldEllipticCurve.x, pointX)
			multiplied, _ = Multiply(s, subtracted)
			pointY, _ := Subtract(multiplied, &thisFiniteFieldEllipticCurve.y)

			return NewFiniteFieldEllipticCurve(pointX, pointY, &thisFiniteFieldEllipticCurve.a, &thisFiniteFieldEllipticCurve.b)

			// 4. 동일한 점인데 접선이 y축에 평행한 경우 y 값은 0이 되어 분모가 0이 되는 에러 발생.
			// 예외처리
		} else {
			infiniteOriginPointX := NewFiniteField(math.MaxInt, thisFiniteFieldEllipticCurve.x.prime)
			infiniteOriginPointY := NewFiniteField(math.MaxInt, thisFiniteFieldEllipticCurve.y.prime)
			return NewFiniteFieldEllipticCurve(infiniteOriginPointX, infiniteOriginPointY, &thisFiniteFieldEllipticCurve.a, &thisFiniteFieldEllipticCurve.b)
		}

		// 2. x1 != x2 인 경우의 덧셈
	} else {

		variationY, _ := Subtract(&bFiniteFieldEllipticCurve.y, &thisFiniteFieldEllipticCurve.y)
		variationX, _ := Subtract(&bFiniteFieldEllipticCurve.x, &thisFiniteFieldEllipticCurve.x)

		//m is inclination of a - b line
		var m, _ = Divide(variationY, variationX)

		// ⚠️ golang 에서 -3 / -110 은 0이다.

		powered, _ := Pow(m, 2)
		subtracted, _ := Subtract(powered, &thisFiniteFieldEllipticCurve.x)
		pointX, _ := Subtract(subtracted, &bFiniteFieldEllipticCurve.x)

		subtracted, _ = Subtract(&thisFiniteFieldEllipticCurve.x, pointX)
		multiplied, _ := Multiply(m, subtracted)
		pointY, _ := Subtract(multiplied, &thisFiniteFieldEllipticCurve.y)
		//fmt.Printf("m:%d, x1 : %d, x3: %d, y1:%d\n", m.number, thisFiniteFieldEllipticCurve.x.number, pointX.number, thisFiniteFieldEllipticCurve.y.number)
		return NewFiniteFieldEllipticCurve(pointX, pointY, &thisFiniteFieldEllipticCurve.a, &thisFiniteFieldEllipticCurve.b)
	}

}

// ScalarMultiply is coefficient of finite field elliptic curve operator
// coefficient : 계수
func (thisFiniteFieldEllipticCurve *finiteFieldEllipticCurve) ScalarMultiply(coefficient int) *finiteFieldEllipticCurve {
	// binary expansion == 이진수 전개
	// Time complexity: log_2(N)

	// It takes only 40 bytes for expressing once trillion in binary number

	infiniteOriginPointX := NewFiniteField(math.MaxInt, thisFiniteFieldEllipticCurve.x.prime)
	infiniteOriginPointY := NewFiniteField(math.MaxInt, thisFiniteFieldEllipticCurve.y.prime)

	current := thisFiniteFieldEllipticCurve

	// point at the infinite origin
	result, _ := NewFiniteFieldEllipticCurve(infiniteOriginPointX, infiniteOriginPointY, &thisFiniteFieldEllipticCurve.a, &thisFiniteFieldEllipticCurve.b)

	for coefficient != 0 {
		// &: bitwise AND
		// Check: LSB is 1?
		if coefficient&1 == 1 {
			result, _ = result.Add(current)
		}

		current, _ = current.Add(current)

		// >>: right shift
		// abandon LSB and check next!
		coefficient >>= 1
	}

	return result

}
