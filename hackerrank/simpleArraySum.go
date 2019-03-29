package hackerrank

func simpleArraySum(ar []int32) int32 {
	sum := int32(0)
	for _, x := range ar {
		sum += x
	}
	return sum
}
