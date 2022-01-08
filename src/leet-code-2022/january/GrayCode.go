package main

func main() {
	code := grayCode(5)
	for i, v := range code {
		println(i)
		println(v)
	}
}

func grayCode(n int) []int {
	ans := make([]int, 1, 1<<n)
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]|1<<(i-1))
		}
	}
	return ans
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/gray-code/solution/ge-lei-bian-ma-by-leetcode-solution-cqi7/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
