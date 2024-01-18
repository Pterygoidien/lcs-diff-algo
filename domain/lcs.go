package domain

// LCS : Longest Common Subsequence
// There are 3 ways to solve this problem:
// 1. Recursive solution
// 2. Dynamic programming
// 3. Dynamic programming with memoization

// lcs is a recursive function that calculates the length of the longest common subsequence (LCS) between two strings, X and Y.
// It takes the strings X and Y, and the lengths of X and Y as input parameters.
// The function returns the length of the LCS.
func Lcs(X, Y string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	} else if X[m-1] == Y[n-1] {
		return 1 + Lcs(X, Y, m-1, n-1)
	} else {
		return max(Lcs(X, Y, m, n-1), Lcs(X, Y, m-1, n))
	}
}

// lcsDP is a dynamic programming function that calculates the length of the longest common subsequence (LCS) between two strings, X and Y.
func LcsDP(X, Y string) int {
	var m, n int = len(X), len(Y)

	//create a 2D slice of size (m+1) x (n+1)
	dp := make([][]int, m+1)

	//initialize the first row and column of the slice to 0
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the matrix
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				// max of the two values above and to the left
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

/*
Video explanation of the algorithm: https://www.youtube.com/watch?v=sSno9rV8Rhg, https://www.youtube.com/watch?v=HgUOWB0StNE

Implementation example with recurvise solution:
X = "AGGTAB"
Y = "GXTXAYB"
m = len(X) = 6
n = len(Y) = 7

LCS should return 4, since the longest common subsequence is GTAB.

So, we start with the last character of each string, and compare them.
If they are equal, we add 1 to the result and move to the next character of each string.
If they are not equal, we move to the next character of each string, and compare them again.
We repeat this process until we reach the end of either string.

In the first iteration, we compare the last character of each string, B and B, at indexes 5 (m-1) and 6 (n-1).
Since they are equal, we add 1 to the result, and move to the next character.

B = B -> +1 to result (1)

So we call LCS again, this time with the same strings but with the lengths reduced by 1, so at
indexes 4 (m-1) and 5 (n-1).In the second iteration, we compare the last character of each string, A and Y, at indexes 4 (m-1) and 5 (n-1).
Since they are not equal, we move to the next character of each string, and compare them again.

A != Y -> So the second statement if also not satisfied, which brings use to the third statement (else), which will return the max
of the two values above and to the left. In this case, the value above is 1, and the value to the left is 0, so we return 1.

So we divide in two threads : one that that evaluates at index 4 (m-1) and 6 (n), and one that evaluates at index 5 (m) and 5 (n-1).
	Thread 1 :
	Thread 2 :






*/
