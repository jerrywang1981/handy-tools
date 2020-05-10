package util

func Min(x ...int) int {
	if len(x) == 0 {
		return 0
	}
	if len(x) == 1 {
		return x[0]
	}
	if len(x) == 2 {
		if x[0] > x[1] {
			return x[1]
		} else {
			return x[0]
		}
	}
	return Min(x[0], Min(x[1:]...))
}

func Max(x ...int) int {
	if len(x) == 0 {
		return 0
	}
	if len(x) == 1 {
		return x[0]
	}
	if len(x) == 2 {
		if x[0] < x[1] {
			return x[1]
		} else {
			return x[0]
		}
	}
	return Max(x[0], Max(x[1:]...))
}
