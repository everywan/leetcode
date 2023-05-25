package codes

import (
	"math/rand"
)

/*

已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。
*/

/*
方法一:
	摇两次骰子, 49个可能. 取四个为一组
	4/49, 每个组内输出值即可. 这里需要注意的是, 71 翻译为十进制就是 8, 72=>9,
	所以 77 可以翻译为 十进制, 然后 % 10 即可, 抛弃 >40 的
方法二:
	第一次取, 只选 [0,4], 每个概率 1/5
	第二次选, <3 1个, >3 一个, 概率为 1/2
	拼出结果 1/10
*/
func rand10() int {
	i := rand7()
	j := rand7()
	if data := i*7 + j; data < 40 {
		return data%10 + 1
	}
	return rand10()
}

func rand7() int {
	return rand.Int() % 7
}

func rand10_2() int {
	for {
		first := rand7()
		if first < 5 {
			for {
				sec := rand7()
				if sec == 3 {
					continue
				}
				if sec < 3 {
					return first
				}
				return first + 5
			}
		}
	}
}
