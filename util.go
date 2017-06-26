package main

func sliceContains(slice []string, element string) bool {
	for _, x := range slice {
		if x == element {
			return true
		}
	}
	return false
}
