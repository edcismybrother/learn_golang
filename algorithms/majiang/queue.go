package main

import "fmt"

func main() {
	// result := hu([]int{1, 1, 1, 2, 3, 3, 3, 4, 4, 5, 5})
	result := hu([]int{4, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4})
	for _, v := range result {
		fmt.Println(v)
	}
}

func hu(pai []int) [][][]int {
	single := [][][]int{}
	ms := make(map[int]bool)
	for i := 0; i < len(pai); i++ {
		target := pai[i]
		if !ms[target] {
			num := 1
			for j := i + 1; j < len(pai); j++ {
				if target == pai[j] {
					num++
				}
			}
			if num > 1 {
				single = append(single, [][]int{[]int{target, target}})
			}
			ms[target] = true
		}
	}
	combNum := len(single)
	for len(single) > 0 {
		if combNum == 0 {
			combNum = len(single)
		}
		first := single[0]
		now := make([]int, len(pai))
		copy(now, pai)
		for _, v := range first {
			for _, v1 := range v {
				for i := 0; i < len(now); i++ {
					v2 := now[i]
					if v1 == v2 {
						now = append(now[:i], now[i+1:]...)
						break
					}
				}
			}
		}
		if len(now) == 0 {
			break
		}
		keNum := 0
		var left1, left2, right1, right2 bool
		for i := 0; i < len(now); i++ {
			if now[i] == now[0] {
				keNum++
			}
			if now[i] == now[0]-2 {
				left1 = true
			}
			if now[i] == now[0]-1 {
				left2 = true
			}
			if now[i] == now[0]+1 {
				right1 = true
			}
			if now[i] == now[0]+2 {
				right2 = true
			}
		}
		if keNum >= 3 {
			single = append(single, append(first, []int{now[0], now[0], now[0]}))
		}
		if left1 && left2 {
			single = append(single, append(first, []int{now[0] - 2, now[0] - 1, now[0]}))
		}
		if left2 && right1 {
			single = append(single, append(first, []int{now[0] - 1, now[0], now[0] + 1}))
		}
		if right1 && right2 {
			single = append(single, append(first, []int{now[0], now[0] + 1, now[0] + 2}))
		}
		combNum--
		single = single[1:]
	}
	return single
}
