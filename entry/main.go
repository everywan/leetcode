package main

import (
	"fmt"

	"github.com/everywan/leetcode/codes"
)

func main() {
	s := new(solutions)
	s.t4()
}

type solutions struct{}

func (s *solutions) t4() {
	codes.FindMedianSortedArrays()
}

func (s *solutions) t5() {
	fmt.Println(codes.LongestPalindrome("cna"))
}

func (s *solutions) t6() {
	str := "012345678"
	numRows := 3
	fmt.Println(codes.Convert(str, numRows))
}

func (s *solutions) t7() {
	fmt.Println(codes.Reverse(123))
}

func (s *solutions) t11() {
	fmt.Println(codes.MaxArea([]int{1, 2, 3, 4, 5}))
}

func (s *solutions) t14() {
	fmt.Println(codes.LongestCommonPrefix([]string{"asd", "awfxc", "asdwec"}))
}

func (s *solutions) t15() {
	fmt.Println(codes.ThreeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))
}

func (s *solutions) t16() {
	fmt.Println(codes.ThreeSumClosest([]int{-1, 2, 1, -4}, 1))
}
