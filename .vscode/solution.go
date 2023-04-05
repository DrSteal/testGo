package main

import (
	"fmt"
	"math"
)

func main() {
	eq := equationDis{}
	eq.a = -2
	eq.b = 5
	eq.c = 5
	eq.formlDis(eq)
}

type equationDis struct {
	a, b, c, D float64
}

func (i *equationDis) formlDis(arg equationDis) {
	i.D = (arg.b * arg.b) - 4*(arg.a*arg.c)

	var x1, x2 float64

	if i.D > 0 {
		x1 = (-float64(i.b) + math.Sqrt(i.D)) / (2 * i.a)
		x2 = (-float64(i.b) - math.Sqrt(i.D)) / (2 * i.a)

		fmt.Println("Your the equation has 2 roots\nD = " + fmt.Sprint(i.D))
		fmt.Println("X1 = " + fmt.Sprint(x1) + "\nX2 = " + fmt.Sprint(x2))
	} else if i.D == 0 {

		var x float64

		x = (-float64(i.b)) / (2 * i.a)

		fmt.Println("Your the equation has 1 roots\nD = " + fmt.Sprint(i.D))
		fmt.Println("X = " + fmt.Sprint(x))

	} else if i.D < 0 {
		fmt.Println("Your the equation hasn't  root\nD = " + fmt.Sprint(i.D))
	}

}
