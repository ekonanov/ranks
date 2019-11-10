package main

/****************************
	challenge: https://www.hackerrank.com/challenges/almost-sorted/problem

	Given an array of integers, determine whether the array can be sorted in ascending order
	using only one of the following operations one time.
		* Swap two elements.
		* Reverse one sub-segment.
	Determine whether one, both or neither of the operations will complete the task. If both
	work, choose swap.

	Output format:
	1. If the array is already sorted, output yes on the first line.
	2. If you can sort this array using one single operation (from the two permitted operations)
		then output yes on the first line and then:
		a. If elements can be swapped, output swap x y in the second line. x and y are the indices
			of the elements to be swapped, assuming that the array is indexed from 1.
		b. Otherwise, when reversing the segment [x..y], output reverse x y in the second line.
			x and y are the indices of the first and last elements of the subsequence to be reversed, both inclusive.
	3. If you cannot sort the array either way, output no on the first line.
****************************/
import (
	"fmt"
)

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func revers(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func isSorted(arr []int, x int) bool {
	lenA := len(arr)
	if x <= 0 {
		x = 1
	}
	for ; x < lenA; x++ {
		if arr[x-1] > arr[x] {
			break
		}
	}
	return x == lenA
}

func almostSorted(arr []int) {
	var (
		lenA      int
		x, y      int
		operation string
	)
	lenA = len(arr)
	for x = 0; x < lenA-1; {
		if arr[x+1] < arr[x] {
			//	нашли нарушение монотонно возрастающей последовательности
			//	1,2,8,4, ...
			//	    ^ мы здесь, следующий элемет меньше текущего
			y = x + 1
			if y == (lenA - 1) {
				//	если у - последний элемент, просто меняем местами, выходим
				swap(arr, x, y)
				operation += fmt.Sprintf("swap %d %d", x+1, y+1)
				break
			}
			asc := (arr[y+1] > arr[y])
			//	1,2,8,4,5,6,7,3,9
			//	    ^ после элемента [х] идет возрастающая последовательность
			//	ищем позицию для [x] такую, что [x]<[y] и выполняем swap(x,y-1)
			if asc {
				maxX := arr[x]
				for ; (y < lenA) && (arr[y] < maxX); y++ {
				}
				y--
				swap(arr, x, y)
				operation += fmt.Sprintf("swap %d %d", x+1, y+1)
			} else {
				//	1,2,8,7,6,5,4,3,9
				//	    ^ после элемента [x] идет убывающая последовательность
				//	ищем ее окончание: [y-1] < [y] и выполняем revers(x,y-1)
				prev := arr[y]
				for ; (y < lenA) && (prev >= arr[y]); y++ {
					prev = arr[y]
				}
				y--
				revers(arr, x, y)
				operation += fmt.Sprintf("reverse %d %d", x+1, y+1)
			}
			break
		}
		x++
	}

	if isSorted(arr, x-1) {
		fmt.Println("yes")
		fmt.Println(operation)
	} else {
		fmt.Println("no")
	}
}

func main() {

	arrs := [][]int{
		{1, 2},
		{2, 1},
		{1, 2, 3},
		{3, 2, 1},
		{3, 1, 2},
		{3, 2, 1, 4},
		{4, 2, 3, 1},
		{4, 3, 2, 1},
		{1, 2, 3, 5, 4},
		{1, 2, 8, 4, 5, 6, 7, 3, 9},
		{1, 2, 8, 4, 5, 3, 6, 7, 9},
		{1, 2, 8, 7, 6, 5, 4, 3, 9, 10},
	}

	for idx, arr := range arrs {
		fmt.Println("test", idx, arr)
		almostSorted(arr)
		fmt.Println("result :", arr)
		fmt.Println()
	}
}
