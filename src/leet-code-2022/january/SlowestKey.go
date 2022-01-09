package main

func main() {
	releaseTimes := make([]int, 2)
	println(slowestKey(append(releaseTimes, 1, 2), "ba"))
}

//func slowestKey(releaseTimes int[], keysPressed string) byte {
//	return keysPressed[0]
//}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	ans := keysPressed[0]
	maxTime := releaseTimes[0]
	for i := 1; i < len(keysPressed); i++ {
		key := keysPressed[i]
		time := releaseTimes[i] - releaseTimes[i-1]
		if time > maxTime || time == maxTime && key > ans {
			ans = key
			maxTime = time
		}
	}
	return ans
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/slowest-key/solution/an-jian-chi-xu-shi-jian-zui-chang-de-jia-yn7u/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
