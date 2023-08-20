package slice

type FilterCallback[Input any, Output any] func(Input, int) Output

func SliceFilter[Input any](input []Input, callback FilterCallback[Input, bool]) []Input {
	var newSlice []Input

	for index, val := range input {
		if callback(val, index) {
			newSlice = append(newSlice, val)
		}
	}

	return newSlice
}
