package insertion

// SortIntSlice - sorted int slice on insertion method (right to left)
func SortIntSlice(s []int) []int {
	for i := len(s) - 2; i >= 0; i-- {
		for j := i; j < len(s)-1 && s[j] > s[j+1]; j++ {
			s[j], s[j+1] = s[j+1], s[j]
			// fmt.Printf("\t\tSWAP:%d <-- %d\n", s[j], s[j+1])
		}
		// fmt.Printf("\t%v\n", s)
	}
	return s
}

// SortIntSliceFromLeft -
func SortIntSliceFromLeft(s []int) []int {
	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0 && s[j] > s[j+1]; j-- {
			s[j], s[j+1] = s[j+1], s[j]
			//fmt.Printf("\t\tSWAP:%d <-- %d\n", s[j+1], s[j])
		}
		//fmt.Printf("\t%v\n", s)
	}
	return s
}

// SortIntSliceFromRight - sorted by right to left
func SortIntSliceFromRight(s []int) []int {
	return SortIntSlice(s)
}
