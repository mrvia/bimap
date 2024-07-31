package bimap

func reverseMap[K, V comparable](forward map[K]V) map[V]K {
	reverse := make(map[V]K)
	// Iterate over the forward map
	for key, value := range forward {
		reverse[value] = key
	}
	return reverse
}

func contains[K comparable](slice []K, value K) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
