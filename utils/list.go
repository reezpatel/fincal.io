package utils

func IncludesInList(list *[]string, item string) bool {
	for _, e := range *list {
		if e == item {
			return true
		}
	}

	return false
}
