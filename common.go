package main

const (
	MinNum = 0
	MaxNum = 9
	N = 4
)

//评分A、B
func score(try, secret [N]int) (a,b int) {
	a = scoreA(try, secret)
	b = scoreAB(try, secret) - a
	return
}

func scoreA(try, secret [N]int) int {
	s := 0
	for i,_ := range try {
		if try[i] == secret[i] {
			s++
		}
	}
	return s
}

func scoreAB(try, secret [N]int) int {
	s := make(map[int]int)
	for _,num := range secret {
		s[num]++
	}

	t := make(map[int]int)
	for _,num := range try {
		t[num]++
	}

	sum := 0
	for num,cnt := range t {
		v,_ := s[num]
		sum += min(cnt, v)
	}

	return sum
}

//全空间遍历
func traverse(repeat bool, f func([N]int)) {
	var v [N]int
	for {
		if repeat || !hasRepeatNum(v) {
			f(v)
		}
		v[0]++
		for i,_ := range v {
			if v[i] <= MaxNum {
				break
			}
			if i == N-1 {
				return
			}
			v[i] = MinNum
			v[i+1]++
		}
	}
}

//获取整个解空间集合
func getAll(repeat bool) [][N]int {
	result := make([][N]int, 0)
	traverse(repeat, func(posible [N]int) {
		result = append(result, posible)
	})
	return result
}

//是否有重复的数字
func hasRepeatNum(arr [N]int) bool {
	for i,_ := range arr {
		for j:=i+1 ; j<len(arr) ; j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

func min(x,y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

