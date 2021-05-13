package day4

import (
	"strconv"
)

func DecodeString(s string) string {
	str, _ := decodeStringHandle(s, 0)
	return str
}
func decodeStringHandle(s string, i int) (string, int) {
	res := "" //字母
	n := ""   //数字
	for i < len(s) {
		tmp := string(s[i])
		//字母
		if tmp >= "a" && tmp <= "z" || tmp >= "A" && tmp <= "Z" {
			res += tmp
			i++
		}
		//数字
		if tmp >= "0" && tmp <= "9" {
			n += tmp
			i++
		}
		//前括号
		if tmp == "[" {
			r := ""
			r, i = decodeStringHandle(s, i+1) //新起始点 返回要重复n次的字符
			tmpn, _ := strconv.Atoi(n)
			for m := 0; m < tmpn; m++ {
				res += r
			}
			n = ""
			continue
		}
		//后括号
		if tmp == "]" {
			return res, i + 1
		}
	}
	return res, 0
}
