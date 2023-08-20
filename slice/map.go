package slice

type MapCallback[Input any, Output any] func(Input, int) Output

func Map[Input any, Output any](input []Input, callback MapCallback[Input, Output]) []Output {
	var newSlice []Output

	for index, val := range input {
		newSlice = append(newSlice, callback(val, index))
	}

	return newSlice
}
