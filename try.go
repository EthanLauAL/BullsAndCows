package main

import (
	"math"
)

const (
	MaxProc = 256
)

//获取最佳（熵期望最大）的猜测（并发执行）
func getBestTry(repeat bool, posible [][N]int) [N]int {
	var result [N]int
	var maxMeanEntropy float64
	type TryResult struct {
		try [N]int
		entropy float64
	}
	channel := make(chan TryResult, MaxProc)
	exit := make(chan int)

	go func(){
		for {
			try_result,ok := <-channel
			if !ok {
				break
			}
			if try_result.entropy > maxMeanEntropy {
				maxMeanEntropy = try_result.entropy
				result = try_result.try
			}
		}
		exit<-1
	}()

	running := make(chan int, MaxProc)
	for i:=0 ; i<MaxProc ; i++ {
		running <- 1
	}
	traverse(repeat, func(try [N]int) {
		<- running
		go func() {
			channel <- TryResult{
				try,
				meanEntropyOf(try, posible),
			}
			running <- 1
		}()
	})
	for i:=0 ; i<MaxProc ; i++ {
		<- running
	}
	close(channel)
	<-exit

	return result
}

//某种猜测的熵期望
func meanEntropyOf(try [N]int, posible [][N]int) float64 {
	stat := make(map[int]int)
	for _,secret := range posible {
		a,b := score(try, secret)
		stat[a<<16 | b]++
	}
	count := len(posible)
	var sumMeanEntropy float64
	for _,s := range stat {
		p := float64(s)/float64(count) //某种输出的可能性
		sumMeanEntropy += p * math.Log2(1/p)
	}
	return sumMeanEntropy
}

//过滤可能性
func filter(try [N]int, posible [][N]int, a, b int) [][N]int {
	result := make([][N]int,0)
	for _,secret := range posible {
		ta,tb := score(try, secret)
		if ta == a && tb == b {
			result = append(result, secret)
		}
	}
	return result
}
