package hash

// works for rising sequences (doesn't check 1/n)
// can be more efficient
func countTriplets(arr []int, r int) int {
	if len(arr) < 3 {
		return 0
	}

	var count int

	valueIndexes := make(map[int][]int)
	valueIndexes[arr[0]] = []int{0}
	valueIndexes[arr[1]] = []int{1}

	for i, v := range arr[2:] {
		valueIndexes[v] = append(valueIndexes[v], i+2)
	}

	for k, indexes := range valueIndexes {
		if k%(r*r) != 0 {
			continue
		}
		prevs := valueIndexes[k/r]
		pprevs := valueIndexes[k/r/r]
		if len(prevs) == 0 || len(pprevs) == 0 {
			continue
		}

		for _, idxMax := range indexes {

			for pi := len(prevs) - 1; pi >= 0; pi-- {
				idxMed := prevs[pi]
				if idxMed > idxMax {
					continue
				}
				var numOfOccurs int
				for ppi := len(pprevs) - 1; ppi >= 0; ppi-- {
					var idxMin = pprevs[ppi]
					if idxMin < idxMed {
						numOfOccurs = ppi + 1
						break
					}
				}

				count += numOfOccurs
			}
		}
	}

	return count
}
