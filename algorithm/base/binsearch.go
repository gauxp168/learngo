package main

func binSearch(arr []int, search int) int {
	low := 0
	high := len(arr)-1
	for low<high {
		mid := (low + high)/2
		if arr[mid] < search {
			low = mid+1
		}else if arr[mid] > search {
			high = mid -1
		}else {
			return mid
		}
	}
	return -1
}
