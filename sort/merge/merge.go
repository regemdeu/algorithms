package merge

import (
	"sync"
)

// SortIntSlice -
func SortIntSlice(s []int) []int {

	if i := len(s) / 2; i > 0 {
		var wg sync.WaitGroup
		wg.Add(2)

		var left []int
		var right []int
		go func() {
			left = SortIntSlice(s[:i])
			wg.Done()
		}()
		go func() {
			right = SortIntSlice(s[i:])
			wg.Done()
		}()
		wg.Wait()
		return merge(left, right)
	}
	return s
}

func merge(left, right []int) []int {
	r := make([]int, len(left)+len(right))

	i, j := 0, 0
	for len(left) > i && len(right) > j {
		if left[i] < right[j] {
			r[i+j] = left[i]
			i++
		} else if len(right) > j {
			r[i+j] = right[j]
			j++
		}
	}

	for len(left) > i {
		r[i+j] = left[i]
		i++
	}

	for len(right) > j {
		r[i+j] = right[j]
		j++
	}

	return r
}
