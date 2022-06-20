package opreturn

func declareReturnTypeOnly() string {
	var allStr string
	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s {
		allStr += str
	}

	return allStr
}

func declareReturnNameTypeWithFmt() (allStr string) {
	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s {
		allStr += str
	}

	return
}

func declareReturnNameTypeWithLog() (allStr string) {
	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s {
		allStr += str
	}

	return
}
