package iteration

func Repeat(char string, times int) string {
	var repetead string

	for i := 0; i < times; i++ {
		repetead += char
	}

	return repetead
}
