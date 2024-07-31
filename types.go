// Package bimap
package bimap

// BiMap
type BiMap[K, V comparable] interface {
	ReadOnlyBiMap[K, V]

	// Put associates the given key with the given value.
	Put(key K, value V)

	// Delete removes the given key from the map.
	Delete(key K)

	// Clear removes all key-value pairs from the map.
	Clear()
}

// ReadOnlyBiMap
type ReadOnlyBiMap[K, V comparable] interface {
	// GetByKey returns the value associated with the given key.
	GetByKey(key K) (V, bool)

	// GetByValue returns the key associated with the given value.
	GetByValue(value V) (K, bool)

	// Len returns the number of key-value pairs in the map.
	Len() int

	// Keys returns a slice of all keys in the map.
	Keys() []K

	// Values returns a slice of all values in the map.
	Values() []V
}
