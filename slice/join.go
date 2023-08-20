package slice

func Join(input []string) string {
	var result string

	for _, val := range input {
		result += val
	}

	return result
}
