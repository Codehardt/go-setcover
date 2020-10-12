package setcover

import (
	"fmt"
	"strconv"
	"strings"
)

func setToString(set []int) string {
	res := make([]string, len(set))
	for i, e := range set {
		res[i] = strconv.Itoa(e)
	}
	return fmt.Sprintf("(%s)", strings.Join(res, " "))
}

func setsToString(sets [][]int) string {
	if len(sets) == 0 {
		return "nil"
	}
	res := make([]string, len(sets))
	for i, set := range sets {
		res[i] = setToString(set)
	}
	return fmt.Sprintf("%s", strings.Join(res, " "))
}

func ExampleGreedyCoverage() {
	var first = true
	gc := func(sets [][]int) {
		if !first {
			//fmt.Println("---")
		}
		first = false
		res := GreedyCoverage(sets)
		fmt.Println("IN ", setsToString(sets))
		fmt.Println("OUT", setsToString(res))
	}
	// Basic Examples
	gc([][]int{})
	gc([][]int{{0}})
	gc([][]int{{0, 1}})
	gc([][]int{{0}, {0}})
	gc([][]int{{0}, {1}})
	// Advanced Examples
	gc([][]int{{0, 2}, {1, 3}, {2, 3}, {0, 1}})
	gc([][]int{{0, 1, 2}, {3, 4, 5}, {0, 1, 2, 3, 4, 5}})
	gc([][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}})
	gc([][]int{{1, 2, 3}, {2, 4}, {3, 4}, {4, 5}})
	gc([][]int{{1, 2}, {2, 3, 4, 5}, {6, 7, 8, 9, 10, 11, 12, 13}, {1, 3, 5, 7, 9, 11, 13}, {2, 4, 6, 8, 10, 12, 13}})
	// Output:
	// IN  nil
	// OUT nil
	// IN  (0)
	// OUT (0)
	// IN  (0 1)
	// OUT (0 1)
	// IN  (0) (0)
	// OUT (0)
	// IN  (0) (1)
	// OUT (0) (1)
	// IN  (0 2) (1 3) (2 3) (0 1)
	// OUT (0 2) (1 3)
	// IN  (0 1 2) (3 4 5) (0 1 2 3 4 5)
	// OUT (0 1 2 3 4 5)
	// IN  (0 1) (2 3) (4 5) (6 7)
	// OUT (0 1) (2 3) (4 5) (6 7)
	// IN  (1 2 3) (2 4) (3 4) (4 5)
	// OUT (1 2 3) (4 5)
}
