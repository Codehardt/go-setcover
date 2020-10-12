package setcover

import "sort"

// set holds the original set elements but also
// a map of elements that are not yet covered in the resulting universe.
type set struct {
	elements          []int
	uncoveredElements map[int]struct{}
}

// newSet generates a new set by initializing the uncovered element map.
func newSet(elements []int) (s set) {
	s.elements = elements
	s.uncoveredElements = make(map[int]struct{})
	for _, element := range elements {
		s.uncoveredElements[element] = struct{}{}
	}
	return
}

// filter removes all elements from the uncovered elements map.
func (s *set) filter(filter map[int]struct{}) {
	for element := range filter {
		delete(s.uncoveredElements, element)
	}
}

// GreedyCoverage returns a minimum set of sets that covers the whole universe by using
// the Greedy Set Coverage Algorithm. The resulting subset will still cover the whole universe but
// its not guaranteed that it's the smallest subset.
func GreedyCoverage(s [][]int) (result [][]int) {
	// convert sets to sets that use above's struct
	sets := make([]set, len(s))
	for i, rawSet := range s {
		sets[i] = newSet(rawSet)
	}
	for {
		// search for the set that covers most elements in the universe
		sort.Slice(sets, func(i, j int) bool {
			if len(sets[i].uncoveredElements) == len(sets[j].uncoveredElements) {
				return len(sets[i].elements) < len(sets[j].elements)
			}
			return len(sets[i].uncoveredElements) > len(sets[j].uncoveredElements)
		})
		if len(sets) == 0 || len(sets[0].uncoveredElements) == 0 {
			// no more sets or elements in sets -> universe is now covered in result
			return
		}
		// add the biggest set to the universe and remove it from the remaining sets
		biggestSet := sets[0]
		result = append(result, biggestSet.elements)
		sets = sets[1:]
		// remove elements of the biggest set from the remaining sets
		for i, set := range sets {
			set.filter(biggestSet.uncoveredElements)
			sets[i] = set
		}
	}
}
