# Code check in Go

Solutions for kata problems in go

## Karate Chop
Problem: [Link](http://codekata.com/kata/kata02-karate-chop/).

Solution: [kata_02](kata_o2.go).
```go
    func binarySearch(arr []int, key int, start int, end int) int {
        if start > end {
            return -1
        }
        middle := start + (end-start)/2
        val := arr[middle]

        if val < key {
            return binarySearch(arr, key, middle+1, end)
        } else if val > key {
            return binarySearch(arr, key, start, end-1)
        } else {
            return middle
        }
    }
```



## Anagrams
Problem: [Link](http://codekata.com/kata/kata06-anagrams/).
Solution: [kata_02](kata_o2.go).
```go
    func binarySearch(arr []int, key int, start int, end int) int {
        if start > end {
            return -1
        }
        middle := start + (end-start)/2
        val := arr[middle]

        if val < key {
            return binarySearch(arr, key, middle+1, end)
        } else if val > key {
            return binarySearch(arr, key, start, end-1)
        } else {
            return middle
        }
    }
```