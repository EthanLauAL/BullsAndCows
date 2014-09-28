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
		var try [N]int
		if i == 1 {
			for i,_ := range try {
				try[i] = i % min(N, MaxNum-MinNum+1) + MinNum
			}
		} else {
			try = getBestTry(posible)
		}
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
	if a == N && b == 0 {
		//fmt.Println("OK")
	} else {
		fmt.Println("Wrong Answer")
	}
}
