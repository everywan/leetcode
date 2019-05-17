package main

import (
	"github.com/everywan/leetcode/codes"
)

func main() {
	s := new(solutions)
	s.t15()
}

type solutions struct{}

func (s *solutions) t4() {
	codes.FindMedianSortedArrays()
}

func (s *solutions) t5() {
	codes.LongestPalindrome()
}

func (s *solutions) t6() {
	codes.Convert()
}

func (s *solutions) t7() {
	codes.Reverse()
}

func (s *solutions) t8() {
	codes.MyAtoi()
}

func (s *solutions) t11() {
	codes.MaxArea()
}

func (s *solutions) t14() {
	codes.LongestCommonPrefix()
}

func (s *solutions) t15() {
	codes.ThreeSum()
}

func (s *solutions) t16() {
	codes.ThreeSumClosest()
}
