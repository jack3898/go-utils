package slice

type SortCallback[Input any] func(Input, Input) int

func Sort[Input any](input []Input, sortFn SortCallback[Input]) []Input {
	var result []Input

	for _, val := range input {
		result = append(result, val)

		for i := len(result) - 1; i > 0; i-- {
			if sortFn(result[i], result[i-1]) < 0 {
				// Swap elements at the current index and the previous index
				result[i], result[i-1] = result[i-1], result[i]
			}
		}
	}

	return result
}
