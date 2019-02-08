package utils

// GetIntersect is the util for checking the matches
func GetIntersect(first []string, second []string) []string {
	result := make([]string, 0)

	for _, firstItem := range first {
		for _, secondItem := range second {
			if firstItem == secondItem {
				result = append(result, firstItem)
			}
		}
	}

	return result
}
