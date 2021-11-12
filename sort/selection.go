package sort

func SelectionSort(in []int) {
	var minIndex int
	for i := 0; i < len(in); i++ {
		minIndex = i
		for j := i + 1; j < len(in); j++ {
			if in[j] < in[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			in[minIndex], in[i] = in[i], in[minIndex]
		}
	}
}
