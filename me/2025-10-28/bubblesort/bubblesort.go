package bubblesort

//
func BubbleUp(arr []int) bool {

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			return true
		}
	}

	return false
}

func BubbleSort(arr []int) {
	for {
		if !BubbleUp(arr) {
			break
		}
	}
}
