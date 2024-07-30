// Package bimap ...
package bimap

// BiMap
type BiMap[K, V comparable] interface {
	// Put associates the given key with the given value.
	Put(key K, value V)

	// GetByKey returns the value associated with the given key.
	GetByKey(key K) (V, bool)

	// GetByValue returns the key associated with the given value.
	GetByValue(value V) (K, bool)

	// Delete removes the given key from the map.
	Delete(key K)

	// Len returns the number of key-value pairs in the map.
	Len() int

	// Clear removes all key-value pairs from the map.
	Clear()

	// Keys returns a slice of all keys in the map.
	Keys() []K

	// Values returns a slice of all values in the map.
	Values() []V
}
