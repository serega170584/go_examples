package main

import (
	"bufio"
	"fmt"
	"os"
)

// LgR8D8k7t8KIprKDTQ7aoo7ed6mhKQwWlFxX
// Q7wQk8rqjaH971SqSQJAMgqYyETo4KmlF4ybf
// Q7wLgR8D8Qkk7t88KIrpqjarHKD971SqSQJTQ7aoAMgoq7eYd6yETmhoK4KmlQwWFlF4xybXf

// LgR8D8k7t8KIprKDTQ7aoo7ed6mhKQwWlFx
// Q7wQk8rqjaH971SqSQJAMgqYyETo4KmlF4yb
// Q7wLgR8D8Qkk7t88KIrpqjarHKD971SqSQJTQ7aoAMgoq7eYd6yETmhoK4KmlQwWFlF4xyb

// LgR8D8k7t8KIprKDTQ7aoo7ed6mhKQwWlF
// Q7wQk8rqjaH971SqSQJAMgqYyETo4KmlF4
// Q7wLgR8D8Qkk7t88KIrpqjarHKD971SqSQJTQ7aoAMgoq7eYd6yETmhoK4KmlQwWFlF4

// LgR8D8k7t8KIprKDTQ7aoo7ed6mhKQw
// Q7wQk8rqjaH971SqSQJAMgqYyETo4Kml
// Q7wLgR8D8Qkk7t88KIrpqjarHKD971SqSQJTQ7aoAMgoq7eYd6yETmhoK4KmlQw

// LgR8D8k7t8KIprKDTQ7aoo7ed6mh
// Q7wQk8rqjaH971SqSQJAMgqYyETo
// Q7wLgR8D8Qkk7t88KIrpqjarHKD971SqSQJTQ7aoAMgoq7eYd6yETmho

//K
//4Km
//K4Km

//

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a := []rune(scanner.Text())

	scanner.Scan()
	b := []rune(scanner.Text())

	scanner.Scan()
	c := []rune(scanner.Text())

	n := len(a)
	m := len(b)

	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true

	for i := 1; i < n+1; i++ {
		if a[i-1] == c[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1; j < m+1; j++ {
		if b[j-1] == c[j-1] {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			k := i + j
			dp[i][j] = (a[i-1] == c[k-1] && dp[i-1][j]) || (b[j-1] == c[k-1] && dp[i][j-1])
		}
	}

	fmt.Println(dp)
}
