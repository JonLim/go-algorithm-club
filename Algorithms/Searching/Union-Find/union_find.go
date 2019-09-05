package unionfind

func UnionFind() {
	index, parent, size := map[int]int{}, []int{}, []int{}

	addSetWith := func(element int) {
		index[element] = len(parent)
		parent = append(parent, len(parent))
		size = append(size, 1)
	}

	var setByIndex func(index int) int
	setByIndex = func(index int) int {
		if parent[index] == index {
			return index
		}

		parent[index] = setByIndex(parent[index])
		return parent[index]
	}

	setOf := func(element int) int {
		if indexOfElement, ok := index[element]; ok {
			return setByIndex(indexOfElement)
		}

		return -1 // Have to guard against this response when calling `setOf()`
	}

	inSameSet := func(firstElement int, secondElement int) bool {
		firstSet, secondSet := setOf(firstElement), setOf(secondElement)
		if firstSet != -1 || secondSet != -1 {
			return firstSet == secondSet
		}

		return false
	}

	unionSetsContaining := func(firstElement int, secondElement int) {
		firstSet, secondSet := setOf(firstElement), setOf(secondElement)
		if firstSet == -1 || secondSet == -1 {
			return
		}

		if firstSet != secondSet {
			if size[firstSet] < size[secondSet] {
				parent[firstSet] = secondSet
				size[secondSet] += size[firstSet]
			} else {
				parent[secondSet] = firstSet
				size[firstSet] += size[secondSet]
			}
		}
	}
}
