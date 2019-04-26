package hackerrank

import "math"

// https://www.hackerrank.com/challenges/the-power-sum/problem

func powerSum(X int32, N int32, candidate int) int32 {
	candidateRes := int32(math.Pow(float64(candidate), float64(N)))
	if candidateRes < X {
		return powerSum(X, N, candidate+1) + powerSum(X-candidateRes, N, candidate+1)
	}
	if candidateRes == X {
		return 1
	}

	return 0

}
