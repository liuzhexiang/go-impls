package sort

func MergeSort(in []int) {
	para := make([]int, len(in))
	merge(in, para, 0, len(in)-1)
}

func merge(in, para []int, begin, end int) {
	if begin >= end {
		return
	}
	mid := ((end - begin) >> 1) + begin
	begin1, end1 := begin, mid
	begin2, end2 := mid+1, end
	merge(in, para, begin1, end1)
	merge(in, para, begin2, end2)
	i := begin
	for begin1 <= end1 || begin2 <= end2 {
		if begin2 > end2 {
			para[i] = in[begin1]
			begin1++
		} else if begin1 > end1 {
			para[i] = in[begin2]
			begin2++
		} else if in[begin1] <= in[begin2] {
			para[i] = in[begin1]
			begin1++
		} else {
			para[i] = in[begin2]
			begin2++
		}
		i++
	}
	copy(in[begin:end+1], para[begin:end+1])
}
