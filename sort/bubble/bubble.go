package bubble

// SortIntSlice - sorted you int slice on bubble method
func SortIntSlice(s []int) []int {
	for {
		b := true

		for i := 1; i < len(s); i++ {
			if s[i] < s[i-1] {
				s[i], s[i-1] = s[i-1], s[i]
				b = false
				//fmt.Printf("\t\tSWAP: %d <-- %d\n", s[i], s[i-1])
			}
		}

		//fmt.Printf("\t%v\n", s)

		if b {
			break
		}
	}

	return s
}
