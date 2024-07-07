// apply quick sort using in-place sorting

/* time complexity 
// Average case O(n log n)
// Worst case O(n^2)
*/

/* space complexity 
//  Average case O(log n)
// Worst case O(n)
*/
package main

import "fmt"

// applying quick sort
func quickSort(arr []int, low, high int){
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}


// applying partion
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Original array:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}
