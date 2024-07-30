// Package bimap ...
package bimap

func reverseMap[K, V comparable](forward map[K]V) map[V]K {
	reverse := make(map[V]K)
	// Iterate over the forward map
	for key, value := range forward {
		reverse[value] = key
	}
	return reverse
}
