package main

import (
	"fmt"
	"sort"
)

func main() {
	// [[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]] 1859
	c := [][]int{
		{259,770},
		{448,54},
		{926,667},
		{184,139},
		{840,118},
		{577,469},
	}
// [[10,20],[30,200],[400,50],[30,20]] 110
//	c := [][]int{
//		{10, 20},
//		{30,200},
//		{400,50},
//		{30,20},
//	}
	fmt.Println("Hello, playground", twoCitySchedCost(c))


}
func twoCitySchedCost(costs [][]int) int {
	fmt.Println(costs)
	sort.SliceStable(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})
	fmt.Println(costs)
	tot1 := 0
	tot2 := 0
	for idx, c := range costs {
		if idx > len(costs)/2 {
			tot1+= c[1]
			tot2+= c[0]
		} else {
			tot1+= c[0]
			tot2+= c[1]
		}
	}
	if tot1 < tot2 {
		return tot1
	} else {
		return tot2
	}
}