package main

import domain "diff_alg/domain"

func main() {
	X := "AGGTAB"
	Y := "GXTXAYB"
	println("Length of LCS is ", domain.Lcs(X, Y, len(X), len(Y)))
	domain.Diff(X, Y)

}
