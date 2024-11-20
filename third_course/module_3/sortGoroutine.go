package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var n int
	fmt.Println("Enter the number of elements:")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	partSize := n / 4
	var wg sync.WaitGroup
	wg.Add(4)

	// Sort each partition in a separate goroutine
	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if i == 3 {
			end = n
		}
		go func(start, end int) {
			defer wg.Done()
			subArray := array[start:end]
			fmt.Printf("Sorting subarray: %v\n", subArray)
			sort.Ints(subArray)
			fmt.Printf("Sorted subarray: %v\n", subArray)
		}(start, end)
	}

	wg.Wait()

	// Merge the sorted subarrays
	sortedArray := mergeSortedArrays(array[:partSize], array[partSize:2*partSize], array[2*partSize:3*partSize], array[3*partSize:])
	fmt.Printf("Sorted array: %v\n", sortedArray)
}

func mergeSortedArrays(a, b, c, d []int) []int {
	result := make([]int, 0, len(a)+len(b)+len(c)+len(d))
	i, j, k, l := 0, 0, 0, 0

	for i < len(a) || j < len(b) || k < len(c) || l < len(d) {
		minVal := int(^uint(0) >> 1) // Max int value
		if i < len(a) && a[i] < minVal {
			minVal = a[i]
		}
		if j < len(b) && b[j] < minVal {
			minVal = b[j]
		}
		if k < len(c) && c[k] < minVal {
			minVal = c[k]
		}
		if l < len(d) && d[l] < minVal {
			minVal = d[l]
		}

		if i < len(a) && a[i] == minVal {
			result = append(result, a[i])
			i++
		} else if j < len(b) && b[j] == minVal {
			result = append(result, b[j])
			j++
		} else if k < len(c) && c[k] == minVal {
			result = append(result, c[k])
			k++
		} else if l < len(d) && d[l] == minVal {
			result = append(result, d[l])
			l++
		}
	}

	return result
}
