package sort

func HeapSort(in []int) {
	for i := len(in) - 1; i > 0; i-- {
		for j := (i - 1) / 2; j >= 0; j-- {
			if j*2+1 <= i && in[j*2+1] > in[j] {
				in[j*2+1], in[j] = in[j], in[j*2+1]
			}
			if j*2+2 <= i && in[j*2+2] > in[j] {
				in[j*2+2], in[j] = in[j], in[j*2+2]
			}
		}
		in[i], in[0] = in[0], in[i]
	}
}
