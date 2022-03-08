package main

import (
	"fmt"
)

func pass_by_var(a int) {
	a++

}

func pass_by_ref(a *int) {
	*a++
}
func main() {
	a := 3

	pass_by_var(a)
	println("pass_by_var", a)
	pass_by_ref(&a)
	println("pass_by_ref", a)
	c := 50
	b := 100
	swap(&c, &b)
	fmt.Println(c, b)

	e, f := swap2(2, 3)
	println(e, f)

}

//不推荐写法
func swap(c, b *int) {
	*c, *b = *b, *c
}

//推荐
func swap2(c, b int) (int, int) {
	return b, c
}
