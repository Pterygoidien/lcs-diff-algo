package domain

import (
	"fmt"
)

func Diff(X, Y string) {
	var m, n int = len(X), len(Y)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the dp matrix (same as in LcsDP function)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Trace back through the matrix to find the differences
	i, j := m, n
	for i > 0 && j > 0 {
		if X[i-1] == Y[j-1] {
			// No difference at this character
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			// Deletion in X
			fmt.Printf("Delete '%c' from X at position %d\n", X[i-1], i-1)
			i--
		} else {
			// Insertion in Y
			fmt.Printf("Insert '%c' to Y at position %d\n", Y[j-1], j-1)
			j--
		}
	}

	// Catch any remaining differences
	for i > 0 {
		fmt.Printf("Delete '%c' from X at position %d\n", X[i-1], i-1)
		i--
	}
	for j > 0 {
		fmt.Printf("Insert '%c' to Y at position %d\n", Y[j-1], j-1)
		j--
	}
}
