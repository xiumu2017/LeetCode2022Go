package main

import "os"

func main() {
	depth := maxDepth("(1+(2*3)+((8)/4))+1")
	println(depth)
	os.Exit(depth)
}

func maxDepth(s string) (ans int) {
	size := 0
	for _, ch := range s {
		if ch == '(' {
			size++
			if size > ans {
				ans = size
			}
		} else if ch == ')' {
			size--
		}
	}
	return
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/maximum-nesting-depth-of-the-parentheses/solution/gua-hao-de-zui-da-qian-tao-shen-du-by-le-av5b/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// Run after the build is not possible The 'main' file has the non-main package or does not contain the 'main' function
