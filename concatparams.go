package piscine

func ConcatParams(args []string) string {
	str := ""

	for i, s := range args {
		if i > 0 {
			str += "\n"
		}
		str += s
	}

	return str
}
