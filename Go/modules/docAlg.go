package modules

const (
	DocBubbleSort = `
Implements the Bubble sorting algorithm,the Bubble sorting algorithm compares 
pairs of adjacent elements in the list and swaps them if they are in the wrong order.
This is repeated until no more swaps are necessary, indicating that the list is sorted.

Complexity:
- Worst case: O(n^2), where n is the number of elements in the list.

Returns:
- List of sorted numbers.
`

	DocSelectionSort = `        
It implements the Sort by Selection algorithm.The Sort by Selection algorithm 
algorithm looks for the smallest element in the list and places it in the correct position.
This is repeated for all the elements in the list, dividing the list into two
parts: the sorted part and the unsorted part.

Complexity:
- Worst case: O(n^2), where n is the number of elements in the list.

Returns:
- List of sorted numbers.
`

	DocInsertionSort = `
Implements the insertion sort algorithm.The insertion sort algorithm builds a one-by-one
sorted list by inserting the unsorted items in the correct position. 
builds a one-by-one sorted list, inserting the unsorted elements in the correct position.
At each iteration, it takes an unsorted element, compares it with the sorted
elements and inserts it in the correct position.

Complexity:
- Worst case: O(n^2), where n is the number of elements in the list.

Returns:
- List of sorted numbers.
`

	DocMergeSort = `
Implements the Merge Sort algorithm. The Merge Sort algorithm uses the 
“divide and conquer” strategy to divide the list into smaller halves.
The “divide and conquer” strategy is used to divide the list into smaller halves,
sort them, and then combine them into a larger sorted list.

This implementation uses two auxiliary functions:
- _mergediv([]int) []int for recursively splitting the list.
- _merge([]int, []int) []int for combining two sorted slices into one.

Complexity:
- Worst case: O(n log n), where n is the number of elements in the list.

Returns:
- List of sorted numbers.
`

	DocQuickSort = `
Quick Sort uses a divide-and-conquer strategy to divide the list into smaller sublists,
sort them, and then combine them into a fully sorted list. It uses a technique called
partitioning to select a pivot element and rearrange the list such that all elements
less than the pivot are moved to the left, and all greater elements to the right.

This implementation is composed of two complementary functions:
- _quicksorthelp(numbers []int, low int, high int): Recursively divides and sorts the sublists.
- _partition(numbers []int, low int, high int) int: Selects the pivot and partitions the list.

Complexity:
- Mean and expected case: O(n log n), where n is the number of elements in the list.
- Worst case: O(n^2), but rarely occurs in practice due to intelligent choice of efficient pivots and partitions. 
efficient pivots and partitions.
        
Returns:
- List of sorted numbers.
`
)
