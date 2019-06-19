package sort

func select_sort(arr *[]string) {
	items := *arr
	length := len(items)
	for i:= 0; i < length; i++ {
		smallestItemRank := i
		for j := i + 1; j < length; j++ {
			if items[smallestItemRank] > items[j] {
				smallestItemRank = j
			}
		}
		swap(arr, i, smallestItemRank)
	}
}

func swap(arr *[]string, rank1 int, rank2 int) {
	if rank1 == rank2 {
		return
	}
	// add the parameter check in the future
	items := *arr
	x := items[rank1]
	items[rank1] = items[rank2]
	items[rank2] = x
}

func sorted(arr *[]string) bool {
	items := *arr
	flag := items[0] < items[1]
	for i := 1; i < len(items) - 1; i++ {
		if items[i] < items[i + 1] != flag {
			return false
		}
	}
	return true
}