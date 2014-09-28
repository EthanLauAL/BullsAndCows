package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	puzzle := NewPuzzle()
	posible := getAll()
	
	for i:=1 ; true ; i++ {
		try := getBestTry(posible)
		a,b := puzzle.Try(try)
		posible = filter(try, posible, a, b)
		fmt.Println(i, "Try:",try, "-", a, "A", b, "B",
			"Remains:", len(posible))
		if len(posible) <= 1 {
			break
		}
	}

	fmt.Println("Got it:", posible[0])

	//测试
	a,b := puzzle.Try(posible[0])
	if a == 4 && b == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("Wrong Answer")
	}
}
