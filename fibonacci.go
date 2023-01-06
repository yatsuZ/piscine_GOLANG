package piscine

func Fibonacci(index int) int {
	if index == 0 {
		return 0
	} else if index < 0 {
		return -1
	} else if index == 1 {
		return 1
	}
	resulta := 0
	mtn := 1
	avant := 0
	for i := index; i > 1; i-- {
		resulta = mtn + avant
		avant = mtn
		mtn = resulta
	}
	return resulta
}
