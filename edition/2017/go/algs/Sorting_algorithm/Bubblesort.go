package main

/*
  Bubble sort
 */

import "fmt"


func main() {
    arr := [5]int{2, 12, 7, 1, 43}
    fmt.Println("Initial array is:", arr)
    fmt.Println("")
    
    tmp := 0
    
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr) - 1; j++ {
            if arr[j] > arr[j + 1] {
                tmp = arr[j]
                arr[j] = arr[j + 1]
                arr[j + 1] = tmp
            }
        }
    }
    
    fmt.Println("Sorted array is: ", arr)
}
