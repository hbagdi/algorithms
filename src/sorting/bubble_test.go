package sorting

import "testing"
import "random"

func TestBubbleSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 10; i++ {
		arr := random.RandomArray(10000)
		BubbleSort(arr)
		if !isIntArraySorted(arr) {
			t.Fatal("Failed")
		}
	}
}
