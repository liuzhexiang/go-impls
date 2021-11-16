package sort

func QuickSort(s []int) {
	quickSort(s, 0, len(s)-1)
}

func quickSort(s []int, begin, end int) {
	if begin >= end {
		return
	}
	// mid := begin + (end-begin)>>1
	// key := s[mid]

	i, j := begin, end
	key := s[i]
	for i < j {
		for i < j && s[j] > key {
			j--
		}
		if i < j {
			s[i] = s[j]
			i++
		}
		for i < j && s[i] < key {
			i++
		}
		if i < j {
			s[j] = s[i]
			j--
		}
	}
	s[i] = key
	quickSort(s, begin, i-1)
	quickSort(s, i+1, end)
}
