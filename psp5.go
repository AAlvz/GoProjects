package main

import( "./proyects/integra"
"fmt"
"math")

func main(){
	p1 := integra.Simpson(10.0, 1.1, 9.0)
	p3 := integra.Simpson(10.0, 2.75, 30.0)
	p2 := integra.Simpson(10.0, 1.1812, 10.0)
	p := math.Abs(p1) - math.Abs(p2)
	e := 0.00001
	seg := 10.0
	fmt.Print(p1, "\n")
	fmt.Print(p2, "\n")
	fmt.Print(p3, "\n")
	fmt.Print(p, "\n")
	for p > e{
		p1 = integra.Simpson(seg, 1.1, 9.0)
		p2 = integra.Simpson(seg*2, 1.1, 9.0)
		p = math.Abs(p1) - math.Abs(p2)
		seg++
		fmt.Print(p1, "\n")
		fmt.Print(p2, "\n")
		fmt.Print(p, "\n")
	}
}
