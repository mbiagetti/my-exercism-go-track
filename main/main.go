package main

import (
"fmt"
	"time"
)

func main() {
//	arr := []int{-2,1,-3,4,-1,2,1,-5,4}
//	arr = []int{-1}
//	arr := []int{7,1,5,3,6,4}
//	fmt.Println("Hello, playground",maxSubArray(arr) )
	//moveZeroes(arr)

//	fmt.Println()
//	fmt.Println(maxProfit(arr))
	start := "2020-04-01 00:00:00 +0000 UTC"
	end := "2021-03-31 00:00:00 +0000 UTC"
	inLayout := "2006-01-02 15:04:05 -0700 MST"
	outLayout := "January 2006"
	s, _ := time.Parse(inLayout, start)
	e, _ := time.Parse(inLayout, end)
	for e.After(s){
		fmt.Println(s.Format(outLayout))
		s = s.AddDate(0, 1, 0)
	}
}

func maxProfit(prices []int) int {
	total := 0

	for {
		i,id1,j,id2 := findMinMaxIdx(prices)
		if id2!=0{
			total += i-j
			fmt.Println(total)
			prices = prices[:(id1+id2)]
			fmt.Println(prices)
		} else {
			break
		}
	}
	return total
}

func findMinMaxIdx(prices []int) (int, int,int, int){
	_,_,i,id1:= MinMax(prices)
	fmt.Println(i,id1,prices[id1:])
	j,id2, _,_:= MinMax(prices[id1:])
	return i,id1,j,id2
}


func MinMax(array []int) (int, int,int, int) {
	var max int = array[0]
	var maxIdx int = 0
	var min int = array[0]
	var minIdx int = 0
	for idx, value := range array {
		if max < value {
			max = value
			maxIdx = idx
		}
		if min > value {
			min = value
			minIdx = idx
		}
	}
	return min, minIdx, max, maxIdx
}

func moveZeroes(nums []int)  {
	out := []int{}
	zero := 0
	for _,j := range nums {
		if j == 0 {
			zero++
		} else {
			out = append(out, j)
		}
	}
	fmt.Println(out)
	nums = out
}
func maxSubArray(nums []int) int {
	maxTotal := nums[0]

	for n := range nums {
		total := 0

		for i := n; i < len(nums); i++ {
				total += nums[i]
				if total > maxTotal {
					maxTotal = total
				}
		}
	}

	return maxTotal
}

func _maxSubArray(nums []int) int {
	maxTotal := 0
	total := 0

	for i := 0; i < len(nums); i++ {
		total = 0
		subArr :=  map[int]bool{}
		for j := i; j < len(nums); j++ {
			num := nums[j]
			if !subArr[num] {
				subArr[num] = true
				total += num
				fmt.Println("arr:", subArr, total)
				if total > maxTotal {
					maxTotal = total
				}
			} else {
				fmt.Println("seq:", subArr, total, num)
				break
			}
		}


	}

	return maxTotal
}
