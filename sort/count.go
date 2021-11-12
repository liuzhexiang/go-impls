package sort

func CountingSort(in []int) {
	min, max := in[0], in[0]
	for i := range in {
		if in[i] < min {
			min = in[i]
		}
		if in[i] > max {
			max = in[i]
		}
	}
	t := make([]int, max-min+1)
	for i := range in {
		t[in[i]-min]++
	}
	for i, j := 0, 0; i < len(t); {
		if t[i] == 0 {
			i++
			continue
		}
		t[i]--
		in[j] = i
		j++
	}
}
