package chapter_2

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

var errInvalidParameter error = errors.New("invalid parameter error")
var errInfiniteValue error = errors.New("invalid parameter (x or y is infinite value)")
var errNotOnSameGraph error = errors.New("not on same graph")

// Regular expression :  y^2 = x^3 + ax^2 + b
// 실수 버전
type ellipticCurve struct {
	x int
	y int
	a int
	b int
}

func NewEllipticCurve(x, y, a, b int) (*ellipticCurve, error) {

	// inbuilt support for basic constants and mathmatical functions.

	// int 32bit 기준
	positiveInfiniteValue := math.MaxInt
	negativeInfiniteValue := math.MinInt

	if x == positiveInfiniteValue || x == negativeInfiniteValue || y == positiveInfiniteValue || y == negativeInfiniteValue {
		return nil, errInfiniteValue
	}

	// 192, 105, 0, 7
	// validate equation
	if math.Pow(float64(y), 2) != math.Pow(float64(x), 3)+float64(a*x)+float64(b) {
		leftValue := math.Pow(float64(y), 2)
		rightValue := math.Pow(float64(x), 3) + float64(0*x) + float64(b)
		fmt.Printf("left: %f, right: %f\n", leftValue, rightValue)

		fmt.Printf("Failed to create ellipticCurve on coordinate (x:%d, y:%d)\n", x, y)
		return nil, errInvalidParameter
	} else {
		fmt.Printf("succeed to create ellipticCurve on coordinate (x:%d, y:%d)\n", x, y)
		return &ellipticCurve{x, y, a, b}, nil
	}
}

func Add(aEllipticCurve, bEllipticCurve *ellipticCurve) (*ellipticCurve, error) {

	// 그래프가 다른경우
	if aEllipticCurve.a != bEllipticCurve.a || aEllipticCurve.b != bEllipticCurve.b {
		return nil, errNotOnSameGraph
		// 1. x축에 수직인 직선 (무한원점)
		// 그냥 무한원점 (x = a a는 상수) 를 x 무한대로 표현함.
	} else if aEllipticCurve.x == math.MaxInt || aEllipticCurve.x == math.MinInt {
		return bEllipticCurve, nil
	} else if bEllipticCurve.x == math.MaxInt || bEllipticCurve.x == math.MinInt {
		return aEllipticCurve, nil

		// 1-1. 역원을 더하는 케이스, x 좌표는 서로 같다.
	} else if aEllipticCurve.x == bEllipticCurve.x && aEllipticCurve.y != bEllipticCurve.y {
		// 무한 원점 반환
		return &ellipticCurve{math.MaxInt, math.MaxInt, aEllipticCurve.a, aEllipticCurve.b}, nil

		// 3. P1 == P2 동일한 점에 대한 덧셈 구하기 (접선의 교점 && x축 대칭)
	} else if aEllipticCurve.x == bEllipticCurve.x && aEllipticCurve.y == bEllipticCurve.y {
		if aEllipticCurve.y != 0 {
			// s is inclination of tangent line

			var s int = (3*int(math.Pow(float64(aEllipticCurve.x), 2)) + aEllipticCurve.a) / 2 * aEllipticCurve.y
			// 여기 뭔가 틀린듯? 10 - 18
			var pointX int = int(math.Pow(float64(s), 2)) - 2*aEllipticCurve.x
			var pointY int = s*(aEllipticCurve.x-pointX) - aEllipticCurve.y

			return NewEllipticCurve(pointX, pointY, aEllipticCurve.a, aEllipticCurve.b)

			// 4. 동일한 점인데 접선이 y축에 평행한 경우 y 값은 0이 되어 분모가 0이 되는 에러 발생.
			// 예외처리
		} else {
			return &ellipticCurve{math.MaxInt, math.MaxInt, aEllipticCurve.a, aEllipticCurve.b}, nil
		}

		// 2. x1 != x2 인 경우의 덧셈
	} else {
		// m is inclination of a - b line
		var m int = (bEllipticCurve.y - aEllipticCurve.y) / (bEllipticCurve.x - aEllipticCurve.x)
		var pointX int = int(math.Pow(float64(m), 2)) - aEllipticCurve.x - bEllipticCurve.x
		var pointY int = m*(aEllipticCurve.x-pointX) - aEllipticCurve.y

		return NewEllipticCurve(pointX, pointY, aEllipticCurve.a, aEllipticCurve.b)
	}

}

func goNewEllipticCurve(x, y, a, b int, waitGroup *sync.WaitGroup) (*ellipticCurve, error) {
	defer waitGroup.Done()
	// validate equation
	if math.Pow(float64(y), 2) != math.Pow(float64(x), 3)+float64(a*x)+float64(b) {
		fmt.Printf("Failed to create ellipticCurve on coordinate (x:%d, y:%d)\n", x, y)
		return nil, errInvalidParameter
	} else {
		fmt.Printf("succeed to create ellipticCurve on coordinate (x:%d, y:%d)\n", x, y)
		return &ellipticCurve{x, y, a, b}, nil
	}
}

func Test2() {

	// 무조건 적으로 goroutine 쓰는게 좋은 것은 아님.
	// waitGroup 만들고 기다리고 ~

	//waitGroup := sync.WaitGroup{}
	//waitGroup.Add(4)
	a := 5
	b := 7
	startTime := time.Now()
	// test time
	// without goroutine

	NewEllipticCurve(2, 4, a, b)
	NewEllipticCurve(-1, -1, a, b)
	NewEllipticCurve(18, 77, a, b)
	NewEllipticCurve(5, 7, a, b)
	// benchmark test time : 31.661µs

	// with goroutine
	//go goNewEllipticCurve(2, 4, a, b, &waitGroup)
	//go goNewEllipticCurve(-1, -1, a, b, &waitGroup)
	//go goNewEllipticCurve(18, 77, a, b, &waitGroup)
	//go goNewEllipticCurve(5, 7, a, b, &waitGroup)
	//waitGroup.Wait()
	endTime := time.Now()
	fmt.Printf("benchmark test time : %v\n", endTime.Sub(startTime))
	// benchmark test time : 167.269µs
	// 오히려 느렸다. 매우 가벼운일은 그냥 하나에서 다 처리하자.
}
