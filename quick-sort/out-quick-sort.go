// apply quick sort using in-place sorting

/* time complexity 
// Average case O(n log n)
// Worst case O(n^2)
*/

/* space complexity 
//  Average case O(nlog n)
// Worst case O(n)
*/

package main

import "fmt"

func quickSort(arr []int) []int{
	// Handling edge cases
	if len(arr) <=1 {
		return arr
	}

	// the main logic
	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}

	for _, v :=range arr{
		if v < pivot{
			left = append(left, v)
		}else if v > pivot{
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

// func quickSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	pivot := arr[len(arr)/2]
// 	left := []int{}
// 	right := []int{}
// 	equal := []int{}

// 	for _, v := range arr {
// 		if v < pivot {
// 			left = append(left, v)
// 		} else if v > pivot {
// 			right = append(right, v)
// 		} else {
// 			equal = append(equal, v)
// 		}
// 	}

// 	left = quickSort(left)
// 	right = quickSort(right)

// 	return append(append(left, equal...), right...)
// }


func main(){
	arr := []int{3,6,8,10,1,2,1}
	fmt.Println("Main array: ", arr)
	sortedArr := quickSort(arr)
	fmt.Println("Sorted array: ", sortedArr)
}
