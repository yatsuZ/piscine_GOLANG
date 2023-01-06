package piscine

func ConcatParams(args []string) string {
	S := ""
	for i := range args {
		S = S + args[i]
		if i != len(args)-1 {
			S = S + "\n"
		}
	}
	return S
}
