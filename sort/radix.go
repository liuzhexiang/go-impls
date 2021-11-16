package sort

func RadixSort(n []int) {
	if len(n) == 0 {
		return
	}
	width := maxBit(n)
	// 创建桶
	buckets := make([][]int, 10)
	// 进行width次排序
	for i := 0; i < width; i++ {
		// 清空桶
		for j := 0; j < 10; j++ {
			buckets[j] = buckets[j][:0]
		}
		// 往桶里扔数据
		for _, v := range n {
			bkIdx := (v / getPow10(i)) % 10
			buckets[bkIdx] = append(buckets[bkIdx], v)
		}
		index := 0
		for j := 0; j < 10; j++ {
			copy(n[index:], buckets[j])
			index += len(buckets[j])
		}
	}
}

func getPow10(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}

// 求一个数组最大值的10进制位数
func maxBit(s []int) int {
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	var n int
	for {
		tmp := max / 10
		if tmp <= 0 {
			n++
			break
		}
		n++
		max = tmp
	}
	return n
}
