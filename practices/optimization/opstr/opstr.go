package opstr

const (
	smallString = "TestForGolangStrings"
	longString  = "TestForGolangStringsTestForGolangStringsTestForGolangStringsTestForGolangStrings"
)

func generateRandomLengthStrings(s string) (data []string) {
	for i := 0; i < 100; i++ {
		data = append(data, s)
	}
	return
}
