# Code check in Go
[![Build Status](https://travis-ci.org/Eslem/GoCodeCheck.svg?branch=develop)](https://travis-ci.org/Eslem/GoCodeCheck)

Solutions for kata problems in go

## Karate Chop
Problem: [Link](http://codekata.com/kata/kata02-karate-chop/).

Solution: [kata_02](kata_02.go)

Test: [kata_02_test](kata_02_test.go)
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
#### Implementations
- Iterative
- Recursive
- Slice Array


## Anagrams
Problem: [Link](http://codekata.com/kata/kata06-anagrams/).

Solution: [kata_06](kata_06.go)

Test: [kata_06_test](kata_06_test.go)

```go
    for _, word := range words {
		sorted := strings.ToLower(SortStringByCharacter(word))
		_, ok := anagrams[sorted]
		if !ok {
			anagrams[sorted] = make([]string, 1)
			anagrams[sorted][0] = word
		} else {
			if len(anagrams[sorted]) == 1 {
				count++
			}
			anagrams[sorted] = append(anagrams[sorted], word)
		}
	}
```
