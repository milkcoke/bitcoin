package chapter_1

import (
	"errors"
	"fmt"
	"math"
)

var errDifferentPrime error = errors.New("two finite field has different prime number each other")

type finiteField struct {
	number int
	prime  uint
}

func NewFiniteField(number int, primeNumber uint) *finiteField {
	// 이거 return 한거 자동으로 Heap 에 할당되는 듯 `New` 키워드 쓴 것 처럼 (js, python, java 처럼)
	// 그렇지 않으면 main 에서 instance 받아서 계속 쓰는게 말이 안됨.
	return &finiteField{number, primeNumber}
}

func (thisFiniteField *finiteField) GetNumber() int {
	return thisFiniteField.number
}

func (thisFiniteField *finiteField) GetPrime() uint {
	return thisFiniteField.prime
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
	fmt.Println(minusedValue)

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

// 여기서 값 유실됨
func Divide(aField, bField *finiteField) (*finiteField, error) {
	//var resultNumber int = aField.number * int(math.Pow(float64(bField.number), float64(aField.prime-2))) % int(aField.prime)
	fmt.Println(float64(aField.number) * math.Pow(float64(bField.number), float64(aField.prime-2)))
	resultNumber := int64(float64(aField.number)*math.Pow(float64(bField.number), float64(aField.prime-2))) % int64(aField.prime)
	fmt.Println(resultNumber)
	return &finiteField{int(resultNumber), aField.prime}, nil
}

func Pow(aField *finiteField, exponent int) (*finiteField, error) {

	// Initialize
	resultNumber := 1
	number := aField.number

	// if number >= exponent update it
	number %= exponent

	for exponent > 0 {
		if exponent&1 == 1 {
			resultNumber = resultNumber * number % int(aField.prime)
		}

		exponent = exponent >> 1
		number = number * number % int(aField.prime)
	}

	return NewFiniteField(resultNumber, aField.prime), nil

}

func TestFiniteField() {

	// https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go
	// golang 의 modulo 연산은 음수를 뱉는다.
	//fmt.Println((-5) % 3)

	// order: 위수
	var order uint = 57
	var order2 uint = 13
	var order3 uint = 31

	a := NewFiniteField(44, order)
	b := NewFiniteField(33, order)
	c := NewFiniteField(9, order)
	d := NewFiniteField(29, order)
	e := NewFiniteField(3, order2)
	f := NewFiniteField(12, order2)
	g := NewFiniteField(10, order2)

	dividend := NewFiniteField(3, order3)
	divisor := NewFiniteField(24, order3)
	result1, _ := Add(a, b)
	result2, _ := Subtract(c, d)
	result3, _ := Multiply(e, f)
	result4, _ := Divide(dividend, divisor)
	fmt.Printf("result1 : %d, result2: %d\n", result1.GetNumber(), result2.GetNumber())
	// %t is bool format in go
	fmt.Printf("result3: %t\n", result3.GetNumber() == g.GetNumber())

	fmt.Printf("result4: %d", result4.GetNumber())
}
