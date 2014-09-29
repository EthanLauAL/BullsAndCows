package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var hard, scanning bool
	for _,s := range os.Args {
		if s == "hard" {
			hard = true
		}
		if s == "scan" {
			scanning = true
		}
	}

	posible := getAll(hard)
	var puzzle Puzzle
	if scanning {
		puzzle = NewScanningPuzzle()
	} else {
		puzzle = NewRandomicPuzzle(posible)
	}
	
	for i:=1 ; true ; i++ {
		var try [N]int
		if i == 1 {
			for i,_ := range try {
				try[i] = i % min(N, MaxNum-MinNum+1) + MinNum
			}
		} else {
			try = getBestTry(hard, posible)
		}
		a,b := puzzle.Try(try)
		posible = filter(try, posible, a, b)
		fmt.Println(i, "Try:",try, "-", a, "A", b, "B",
			"Remains:", len(posible))
		if len(posible) <= 1 {
			break
		}
		if len(posible) <= 8 {
			fmt.Println("Remains:",posible)
		}
	}

	if len(posible) == 1 {
		fmt.Println("Got it:", posible[0])
	} else {
		fmt.Println("Somthing wrong...")
	}
}
