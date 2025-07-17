package main

func compareSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range len(a) {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func compare2Dslices[T comparable](a, b [][]T) bool {
	n, m := len(a), len(a[0])

	for i := range n {
		for j := range m {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
