package main

import "fmt"

func main() {
	fmt.Println("neovim-go")

	MaxLongHuiWen("abccccdd")
}

// MaxLongHuiWen ... 最长回文字符串
// abccccdd = 7  得到: dccaccd
//思路:
//    1、先取任意取一个字符，然后取剩余的进行两头拼接，统计长度,时间复杂度很高。
//       1.1 先得到所有的奇数个字符的slice（就是在这个slice里面只出现过一次的，那么剩下的基本就是偶数个字符了）
//       1.2 依次循环任意字符作为回文中心串进行拼接，这种效率最低。
//    2、取重复字符的数量然后取一个奇数的字符就是最长的
//       重点是取有多少个偶数对的字符，然后任意取一个奇数个字符的。
func MaxLongHuiWen(str string) {

}

// TwoNumsAdd ...
// 两数之和
// 给定一个数组，和一个目标值，请你在数组里面找到和为目标的那两个整数，并返回它的下标
// nums = [2,7,11,15] target = 7
// 思路
//
func TwoNumsAdd() {

}

// MonotonicSequence ...
// 单调数列
// array [6,5,2,2] true
func MonotonicSequence() {

}

// ArrayRepeat ...
// 给定一个数组，判断是否有重复元素,如果没有一个重复的就返回false否则就返回true
// 思路:
//      1、从头进行开始循环挨个判断,效率比较低。
//      2、从头开始循环，但是第一个跟第二个，第二个跟第一个，第三个跟第一个，依次类推。有重复的就返回false

func ArrayRepeat() {

}

// RansomLetter ...
// a b false
// aa ab false
// aa aad true
// 思路:
//      1、
func RansomLetter() {

}

// VirtualStack  ...
// 模拟栈 实现push、pop 、top、getMin等方法
func VirtualStack() {

}

// SplicingStr ...
// 给定一个字符串，在指定的字符串里面拼接，返回最多可以拼几个。
func SplicingStr() {

}

// DelSliceRune ... delete slice
func DelSliceRune(src []interface{}, index int) []interface{} {
	return append(src[:index], src[index+1:]...)
}
