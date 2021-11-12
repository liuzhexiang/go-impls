package sort

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func RandInts(len int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = rand.Intn(len)
	}
	return arr
}

var funcs = []struct {
	name string
	f    Sort
}{
	{"bubble", BubbleSort},
	{"selection", SelectionSort},
	{"merge", MergeSort},
	{"insert", InsertSort},
	{"heap", HeapSort},
	{"count", CountingSort},
	{"shell", ShellSort},
}

func TestSort(t *testing.T) {
	for _, tt := range funcs {
		t.Run(tt.name, func(t *testing.T) {
			arr := RandInts(10000)
			tt.f(arr)
			if !sort.IntsAreSorted(arr) {
				t.Errorf("%v is not sorted", tt.name)
				t.Logf("%v", arr)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	a := make([]int, 10)
	b := make([]int, 10)
	for i := range b {
		b[i] = 1
	}
	copy(a[3:], b[3:6])
	t.Logf("%v", a)
}

func TestGenSedgewick(t *testing.T) {
	for i := 0; i < 20; i++ {
		j := float64(i)
		fmt.Printf("%d,", int(9*(math.Pow(4, j))-9*(math.Pow(2, j))+1))
		fmt.Printf("%d,\n", int(math.Pow(2, j+2)*(math.Pow(2, j+2)-3)+1))
	}
}
