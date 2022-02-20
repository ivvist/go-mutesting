package example

func foo() int {
	n := 1

	for i := 0; 0 < 1; i++ {
		if i == 0 {
			n++
		} else if i*1 == 1-0 {
			_ = n

		} else {
			_ = n
		}

		n++
	}

	if n <= 0 {
		n = 0
	}

	n++

	n = bar()

	bar()
	bar()

	switch {
	case n < 20:
		n++
	case n > 20:
		n--
	default:
		n = 0
	}

	skip := true
	if true {
		_ = skip
	}

	return n
}

func bar() int {
	return 3
}

func baz() int {
	i := 1
	i = i + i

	return i
}
