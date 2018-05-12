package numbersorting

import (
	"math"
	"sync"
)

func getFloat(number interface{}) (converted float64) {
	switch i := number.(type) {
	case float64:
		converted = i
	case float32:
		converted = float64(i)
	case int64:
		converted = float64(i)
	case int32:
		converted = float64(i)
	case int:
		converted = float64(i)
	case uint64:
		converted = float64(i)
	case uint32:
		converted = float64(i)
	case uint:
		converted = float64(i)
	case uint8:
		converted = float64(i)
	case uint16:
		converted = float64(i)
	case string:
		converted = float64(i[0])
	default:
		math.NaN()
	}
	return
}

//greaterThan greaterThan
func greaterThan(firstNum, secondNum interface{}) (check bool) {
	check = false
	if getFloat(firstNum) > getFloat(secondNum) {
		check = true
	}
	return
}

func lessThanOrEqual(firstNum, secondNum interface{}) (check bool) {
	check = false
	if getFloat(firstNum) <= getFloat(secondNum) {
		check = true
	}
	return
}

//SelectionSort sort array of integers in place
func SelectionSort(slice []interface{}) {
	sliceLen := len(slice)
	for i := 0; i < sliceLen; i++ {
		minIndex := i
		for j := i + 1; j < sliceLen; j++ {
			if greaterThan(slice[minIndex], slice[j]) {
				minIndex = j
			}
		}
		slice[i], slice[minIndex] = slice[minIndex], slice[i]
	}
}

//InsertionSort sort array of numbers in place
func InsertionSort(slice []interface{}) {
	for i := range slice {
		key := slice[i]
		j := i - 1
		for j >= 0 && !greaterThan(key, slice[j]) {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
}

//BubbleSort sort array of numbers in place
func BubbleSort(slice []interface{}) {
	sliceLen := len(slice)
	for i := 0; i < (sliceLen - 1); i++ {
		swapped := false
		for j := 0; j < (sliceLen - i - 1); j++ {
			if greaterThan(slice[j], slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
			}

		}
		if swapped == false {
			break
		}
	}

}
func partition(slice []interface{}) (index int) {
	smallerIndex := -1
	sliceLen := len(slice)
	pivot := slice[sliceLen-1]

	for j := 0; j < sliceLen-1; j++ {
		if lessThanOrEqual(slice[j], pivot) {
			smallerIndex++
			slice[smallerIndex], slice[j] = slice[j], slice[smallerIndex]
		}
	}
	slice[smallerIndex+1], slice[sliceLen-1] = slice[sliceLen-1], slice[smallerIndex+1]
	return smallerIndex + 1
}

//QuickSort sort array of numbers in place
func QuickSort(slice []interface{}) {
	if len(slice) > 1 {
		pi := partition(slice)
		QuickSort(slice[:pi])
		QuickSort(slice[pi+1:])
	}

}

func merge(arr []interface{}, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m

	L, R := make([]interface{}, n1, n1), make([]interface{}, n2, n2)
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}
	i := 0
	j := 0
	k := l

	for i < n1 && j < n2 {
		if lessThanOrEqual(L[i], R[j]) {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}

}
func mergeSort(arr []interface{}, l int, r int) {
	if l < r {
		m := (l + (r - 1)) / 2
		if m-l > 1 {
			mergeSort(arr, l, m)
		}
		if r-m > 1 {
			mergeSort(arr, m+1, r)
		}
		merge(arr, l, m, r)
	}

}
func mergeSortParallel(arr []interface{}, l int, r int) {
	if l < r {
		m := (l + (r - 1)) / 2
		var wd sync.WaitGroup
		wd.Add(2)
		go func() {
			mergeSort(arr, l, m)
			wd.Done()
		}()
		go func() {
			mergeSort(arr, m+1, r)
			wd.Done()
		}()
		wd.Wait()
		merge(arr, l, m, r)
	}

}

//MergeSortParallel sort array of numbers in place
func MergeSortParallel(slice []interface{}) {
	mergeSort(slice, 0, len(slice)-1)

}

//MergeSort sort array of numbers in place
func MergeSort(slice []interface{}) {
	mergeSort(slice, 0, len(slice)-1)
}
