package slice

func SliceJoin(input []string) string {
	var result string

	for _, val := range input {
		result += val
	}

	return result
}
