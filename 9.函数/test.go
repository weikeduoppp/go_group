package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func test() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println("distribution：", distribution)
}

func dispatchCoin() int {
	for _, i := range users {
		_, ok := distribution[i]
		if !ok {
			distribution[i] = 0
		}
		rule := matchRule(i)
		distribution[i] = rule
	}
	return coins
}

func matchRule(s string) int {
	var result int
	for _, i := range s {
		// fmt.Println('e') // rune
		switch {
		case i == 'e' || i == 'E':
			result++
		case i == 'i' || i == 'I':
			result += 2
		case i == 'o' || i == 'O':
			result += 3
		case i == 'u' || i == 'U':
			result += 4
		}
	}
	coins -= result
	return result
}
