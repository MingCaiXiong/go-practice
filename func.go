package main

import (
	"fmt"
	"math"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil

	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

//返回多个值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//函数式编程方式
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

//x的y次方
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))

}
func main() {

	if result, err := eval(6, 3, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		println(result)
	}

	if result, err := eval(6, 3, "*"); err != nil {
		fmt.Println("Error:", err)
	} else {
		println(result)
	}

	println("------------------")
	q, r := div(10, 3)
	fmt.Println(q, r)
	println("---------函数式编程---------")

	fmt.Println(apply(pow, 2, 3))
	println("---------匿名函数---------")
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	println("---------可变参数列表---------")

	println(sum(1, 2, 3, 4, 5))

}
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
