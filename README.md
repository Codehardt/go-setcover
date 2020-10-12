## go-setcover

This Golang package calculates the smallest combination of sets that covers all elements in that sets. This is also known as the [Set cover problem](https://en.wikipedia.org/wiki/Set_cover_problem).

### Example: 

Consider there are four sets `{1, 2, 3}`, `{2, 4}`, `{3, 4}` and `{4, 5}`. The smallest possible combination of those sets that still cover all elements is `{1, 2, 3}`, `{4, 5}`.

```golang
mySets := [][]int{{1, 2, 3}, {2, 4}, {3, 4}, {4, 5}}
minimalIndices := setcover.GreedyCoverageIndex(mySets) // equals []int{0, 3}
```

### Exceptions:

It isn't guaranteed that this package returns the smallest combination of sets because of the usage of the greedy algorithm.

Consider there are five sets `{1, 2}`, `{2, 3, 4, 5}`, `{6, 7, 8, 9, 10, 11, 12, 13}`, `{1, 3, 5, 7, 9, 11, 13}` and `{2, 4, 6, 8, 10, 12, 13}`.
The optimal combination is `{1, 3, 5, 7, 11, 13}` and `{2, 4, 6, 8, 10, 12, 13}` but the greedy algorithm produces `{6, 7, 8, 9, 10, 11, 12, 13}`, `{2, 3, 4, 5}`, `{1, 2}`, 
because it always searches for the sets with the biggest count of uncovered elements, first.