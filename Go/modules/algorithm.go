package modules

import (
	"fmt"

	"github.com/gookit/color"
)

type SortAlg struct {
	numbers []int
	count   int
}

func (self *SortAlg) BubbleSort() []int {
	fmt.Println(DocBubbleSort)
	for i := 0; i < self.count; i++ {
		fmt.Printf("\r[%v] %v %d data: Element -> [%v]", color.Green.Render("+"), color.Yellow.Render("BubbleSort"), self.count, color.Red.Render(self.numbers[i]))
		for j := 0; j < self.count-i-1; j++ {
			if self.numbers[j] > self.numbers[j+1] {
				self.numbers[j], self.numbers[j+1] = self.numbers[j+1], self.numbers[j]
			}
		}
	}
	fmt.Println()
	return self.numbers
}

func (self *SortAlg) SelectionSort() []int {
	fmt.Println(DocSelectionSort)
	for i := 0; i < self.count; i++ {
		fmt.Printf("\r[%v] %v %d data: Element -> [%v]", color.Green.Render("+"), color.Yellow.Render("SelectionSort"), self.count, color.Red.Render(self.numbers[i]))
		mini := i
		for j := i + 1; j < self.count; j++ {
			if self.numbers[j] < self.numbers[mini] {
				mini = j
			}
		}
		self.numbers[i], self.numbers[mini] = self.numbers[mini], self.numbers[i]
	}
	fmt.Println()
	return self.numbers
}

func (self *SortAlg) InsertionSort() []int {
	fmt.Println(DocInsertionSort)
	for i := 0; i < self.count; i++ {
		fmt.Printf("\r[%v] %v %d data: Element -> [%v]", color.Green.Render("+"), color.Yellow.Render("InsertionSort"), self.count, color.Red.Render(self.numbers[i]))
		pos := i
		aux := self.numbers[i]
		for pos > 0 && self.numbers[pos-1] > aux {
			self.numbers[pos] = self.numbers[pos-1]
			pos--
		}
		self.numbers[pos] = aux
	}
	fmt.Println()
	return self.numbers
}

func (self *SortAlg) MergeSort() []int {
	fmt.Println(DocMergeSort)
	self.numbers = self._mergediv(self.numbers)
	fmt.Println()
	return self.numbers
}

func (self *SortAlg) _mergediv(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	mid := len(numbers) / 2
	left := self._mergediv(numbers[:mid])
	right := self._mergediv(numbers[mid:])

	return self._merge(left, right)
}

func (self *SortAlg) _merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			fmt.Printf("\r[%v] %v %d data: Element -> [%v]", color.Green.Render("+"), color.Yellow.Render("MergeSort"), self.count, color.Red.Render(left[i]))
			i++
		} else {
			fmt.Printf("\r[%v] %v %d data: Element -> [%v]", color.Green.Render("+"), color.Yellow.Render("MergeSort"), self.count, color.Red.Render(right[j]))
			result = append(result, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		result = append(result, left[i])
		fmt.Printf("\r[%v] %v %d data: Element -> [%v]", color.Green.Render("+"), color.Yellow.Render("MergeSort"), self.count, color.Red.Render(left[i]))
	}

	for ; j < len(right); j++ {
		result = append(result, right[j])
		fmt.Printf("\r[%v] %v %d data: Element -> [%v]", color.Green.Render("+"), color.Yellow.Render("MergeSort"), self.count, color.Red.Render(right[j]))
	}
	return result
}

func (self *SortAlg) QuickSort() []int {
	fmt.Println(DocQuickSort)
	self._quicksorthelp(self.numbers, 0, self.count-1)
	fmt.Println()
	return self.numbers
}

func (self *SortAlg) _quicksorthelp(numbers []int, low, high int) {
	if low < high {
		pi := self._partition(numbers, low, high)
		self._quicksorthelp(numbers, low, pi-1)
		self._quicksorthelp(numbers, pi+1, high)
	}
}

func (self *SortAlg) _partition(numbers []int, low, high int) int {
	pivot := numbers[high]
	i := low - 1
	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
			fmt.Printf("\r[%v] %v %d data: - Element [%v]", color.Green.Render("+"), color.Yellow.Render("QuickSort"), self.count, color.Red.Render(numbers[i]))
		}
	}
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}
