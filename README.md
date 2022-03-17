# go-list-functions

A small package to demonstrate go generics. This package implements common list operations
* map 
* fold
* reduce 
* contains

(contrived) example:
```
// Fold a list to determine if it contains an even number
source := []int{1, 2, 3, 4}
res := listfns.Fold(source, false, func(acc bool, current int) bool {
    return acc || current % 2 == 0
})
```
