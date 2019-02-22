package utils

// GetIntersect is the util for checking the matches
func GetIntersect(first []string, second []string) []string {
	result := make([]string, 0)

	start := 0
	for _, firstItem := range first {
		for i := start; i< len(second); i++ {
			if firstItem == second[i] {
				start = i
				result = append(result, firstItem)
			}
		}
	}

	return result
}
