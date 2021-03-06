package problem1143

// longestCommonSubsequence two dimensional
func longestCommonSubsequence(text1, text2 string) int {
	// init dp table
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// longestCommonSubsequenceV2 one dimensional
func longestCommonSubsequenceV2(text1, text2 string) int {
	// init dp table
	dp := make([]int, len(text2)+1)

	for i := 1; i <= len(text1); i++ {
		last := 0
		for j := 1; j <= len(text2); j++ {
			tmp := dp[j] // tmp 记录的是二维dp矩阵正上方的值
			if text1[i-1] == text2[j-1] {
				dp[j] = last + 1 // last 记录的是二维dp矩阵左上方的值
			} else {
				dp[j] = max(tmp, dp[j-1])
			}
			last = tmp
		}
	}

	return dp[len(text2)]
}
