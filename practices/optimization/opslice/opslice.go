package opslice

var opStr = [5]string{"one", "two", "3", "4", "five"}

func declareReturnSliceNoLength(s [5]string) (allSlice []string) {
	for sum := 0; sum < 100; sum++ {
		for _, v := range s {
			allSlice = append(allSlice, v)
		}
	}

	return
}

func declareReturnSliceWithLength(s [5]string) (allSlice [500]string) {
	for sum := 0; sum < 100; sum++ {
		for i, v := range s {
			allSlice[i] = v
		}
	}

	return
}
