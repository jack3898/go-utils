package slice

type ReduceCallback[Accumulator any, Input any] func(Accumulator, Input) Accumulator

func SliceReduce[Input any, Output any](input []Input, reducerFn ReduceCallback[Output, Input]) Output {
	var result Output

	for _, val := range input {
		result = reducerFn(result, val)
	}

	return result
}
