package setcover

import (
	"fmt"
	"strconv"
	"strings"
)

func setToString(set []int, index bool) string {
	if len(set) == 0 {
		return "nil"
	}
	res := make([]string, len(set))
	for i, e := range set {
		if index {
			res[i] = "S" + strconv.Itoa(e+1)
		} else {
			res[i] = strconv.Itoa(e)
		}
	}
	return strings.Join(res, " ")
}

func setsToString(sets [][]int, resultset bool) string {
	if len(sets) == 0 {
		if resultset {
			return "nil"
		}
		return "nil\n"
	}
	res := make([]string, len(sets))
	for i, set := range sets {
		if resultset {
			res[i] = fmt.Sprintf("%s", setToString(set, false))
		} else {
			res[i] = fmt.Sprintf("S%d: %s\n", i+1, setToString(set, false))
		}
	}
	sep := ""
	if resultset {
		sep = " # "
	}
	return strings.Join(res, sep)
}

func ExampleGreedyCoverageIndex() {
	gc := func(sets [][]int) {
		res := GreedyCoverageIndex(sets)
		fmt.Print(setsToString(sets, false))
		fmt.Println("Result:", setToString(res, true))
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
	// nil
	// Result: nil
	// S1: 0
	// Result: S1
	// S1: 0 1
	// Result: S1
	// S1: 0
	// S2: 0
	// Result: S1
	// S1: 0
	// S2: 1
	// Result: S1 S2
	// S1: 0 2
	// S2: 1 3
	// S3: 2 3
	// S4: 0 1
	// Result: S1 S2
	// S1: 0 1 2
	// S2: 3 4 5
	// S3: 0 1 2 3 4 5
	// Result: S3
	// S1: 0 1
	// S2: 2 3
	// S3: 4 5
	// S4: 6 7
	// Result: S1 S2 S3 S4
	// S1: 1 2 3
	// S2: 2 4
	// S3: 3 4
	// S4: 4 5
	// Result: S1 S4
	// S1: 1 2
	// S2: 2 3 4 5
	// S3: 6 7 8 9 10 11 12 13
	// S4: 1 3 5 7 9 11 13
	// S5: 2 4 6 8 10 12 13
	// Result: S3 S2 S1
}

func ExampleGreedyCoverage() {
	gc := func(sets [][]int) {
		res := GreedyCoverage(sets)
		fmt.Print(setsToString(sets, false))
		fmt.Println("Result:", setsToString(res, true))
	}
	gc([][]int{})
	gc([][]int{{0}})
	gc([][]int{{0}, {1}})
	gc([][]int{{1, 2}, {2, 3, 4, 5}, {6, 7, 8, 9, 10, 11, 12, 13}, {1, 3, 5, 7, 9, 11, 13}, {2, 4, 6, 8, 10, 12, 13}})
	// Output:
	// nil
	// Result: nil
	// S1: 0
	// Result: 0
	// S1: 0
	// S2: 1
	// Result: 0 # 1
	// S1: 1 2
	// S2: 2 3 4 5
	// S3: 6 7 8 9 10 11 12 13
	// S4: 1 3 5 7 9 11 13
	// S5: 2 4 6 8 10 12 13
	// Result: 6 7 8 9 10 11 12 13 # 2 3 4 5 # 1 2
}
