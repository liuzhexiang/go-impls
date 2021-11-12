package sort

func InsertSort(in []int) {
	for i := 1; i < len(in); i++ {
		for j := 0; j < i; j++ {
			if in[i] <= in[j] {
				tmp := in[i]
				copy(in[j+1:i+1], in[j:i])
				in[j] = tmp
				break
			}
		}
	}
}
