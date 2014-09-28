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
func traverse(f func([N]int)) {
	var v [N]int
	for {
		f(v)
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
func getAll() [][N]int {
	result := make([][N]int, 0)
	traverse(func(posible [N]int) {
		result = append(result, posible)
	})
	return result
}

func min(x,y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

