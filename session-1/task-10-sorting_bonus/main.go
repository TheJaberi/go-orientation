package main

import "fmt"

// BubbleSort implementation
func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        // Flag to optimize bubble sort
        swapped := false
        
        // Last i elements are already in place
        for j := 0; j < n-i-1; j++ {
            // Compare adjacent elements
            if arr[j] > arr[j+1] {
                // Swap them if they are in wrong order
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = true
            }
        }
        
        // If no swapping occurred, array is sorted
        if !swapped {
            break
        }
    }
}

// SelectionSort implementation
func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        // Find the minimum element in unsorted array
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        
        // Swap the found minimum element with the first element
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}

// InsertionSort implementation
func insertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        
        // Move elements of arr[0..i-1] that are greater than key
        // to one position ahead of their current position
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }
}

// Helper function to print array with a label
func printArray(label string, arr []int) {
    fmt.Printf("%s: %v\n", label, arr)
}

// Helper function to make a copy of an array
func copyArray(arr []int) []int {
    newArr := make([]int, len(arr))
    copy(newArr, arr)
    return newArr
}

func main() {
    // Original array
    original := []int{64, 34, 25, 12, 22, 11, 90}
    
    // Bubble Sort
    fmt.Println("\nBubble Sort Demo:")
    bubbleArr := copyArray(original)
    printArray("Original", bubbleArr)
    bubbleSort(bubbleArr)
    printArray("Sorted", bubbleArr)
    
    // Selection Sort
    fmt.Println("\nSelection Sort Demo:")
    selectionArr := copyArray(original)
    printArray("Original", selectionArr)
    selectionSort(selectionArr)
    printArray("Sorted", selectionArr)
    
    // Insertion Sort
    fmt.Println("\nInsertion Sort Demo:")
    insertionArr := copyArray(original)
    printArray("Original", insertionArr)
    insertionSort(insertionArr)
    printArray("Sorted", insertionArr)
}

/* 
Tasks:
1. Implement the bubble sort algorithm
2. Implement another sorting algorithm (selection sort or insertion sort)
3. Bonus: Compare the performance of different sorting algorithms

Concepts covered:
- Sorting algorithms and their implementation
- Array manipulation
- Nested loops
- Algorithm complexity
- In-place vs out-of-place sorting
- Stable vs unstable sorting
- Performance considerations
- Array copying and modification
*/
