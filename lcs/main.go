package main

// 1 0 0 0 0 0 0 0 0 0 0
// 1 0 1 0 0 0 0 0 0 0 0
// 1 1 1 0 0 0 0 0 0 0 0
// 0 0 0 0 0 0 0 0 0 0 0
// 0 0 0 1 1 1 1 1 1 1 1
func main() {
	a := []rune("abcdef")
	b := []rune("bdfg")
	al := len(a)
	bl := len(b)
	dp := make([][]int, al)
	for i := 0; i < al; i++ {
		dp[i] = make([]int, bl)
	}

	if a[0] == b[0] {
		dp[0][0] = 1
	}

	for i := 1; i < al; i++ {
		if a[i] == b[0] {
			dp[i][0] = 1
			continue
		}

		dp[i][0] = dp[i-1][0]
	}

	for i := 1; i < bl; i++ {
		if b[i] == a[0] {
			dp[0][i] = 1
			continue
		}

		dp[0][i] = dp[0][i-1]
	}

	for i := 1; i < al; i++ {
		for j := 1; j < bl; j++ {
			if a[i] == b[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				continue
			}

			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
		}
	}
}
