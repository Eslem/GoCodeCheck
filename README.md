# Code check in Go
[![Build Status](https://travis-ci.org/Eslem/GoCodeCheck.svg?branch=master)](https://travis-ci.org/Eslem/GoCodeCheck)

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
    for _, word := range words {
		sorted := strings.ToLower(SortStringByCharacter(word))
		_, ok := anagrams[sorted]
		if !ok {
			anagrams[sorted] = make([]string, 1)
			anagrams[sorted][0] = word
		} else {
			anagrams[sorted] = append(anagrams[sorted], word)
			count++
        }
    }
```