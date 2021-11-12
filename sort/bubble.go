package sort

func BubbleSort(input []int) {
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
}
