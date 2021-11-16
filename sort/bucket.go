package sort

func BucketSort(s []int) {
	if len(s) == 0 {
		return
	}
	// 找到max和min
	min, max := s[0], s[0]
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
		if s[i] > max {
			max = s[i]
		}
	}
	if min == max {
		return
	}
	// 计算桶的数量
	bucketNum := (max-min)/len(s) + 1
	// 创建桶
	buckets := make([][]int, bucketNum)

	// 遍历原数组，并赋值
	for i := range s {
		nthBucket := (s[i] - min) / len(s)
		buckets[nthBucket] = append(buckets[nthBucket], s[i])
	}
	// 每个桶单独排序
	for _, v := range buckets {
		InsertSort(v)
	}
	// 把每个桶按顺序复制回原数组中
	index := 0
	for _, v := range buckets {
		copy(s[index:], v)
		index += len(v)
	}
}
