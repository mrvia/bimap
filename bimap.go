// Package bimap ...
package bimap

type bimap[K, V comparable] struct {
	forward  map[K]V
	reversed map[V]K
}

// NewBiMap creates a new BiMap instance from a given map.
func NewBiMap[K, V comparable](initMap map[K]V) BiMap[K, V] {
	return &bimap[K, V]{
		forward:  initMap,
		reversed: reverseMap(initMap),
	}
}

// Put adds a new key-value pair to the BiMap.
func (b *bimap[K, V]) Put(key K, value V) {
	b.forward[key] = value
	b.reversed[value] = key
}

// GetByKey returns the value associated with the given key.
func (b *bimap[K, V]) GetByKey(key K) (V, bool) {
	value, ok := b.forward[key]
	return value, ok
}

// GetByValue returns the key associated with the given value.
func (b *bimap[K, V]) GetByValue(value V) (K, bool) {
	key, ok := b.reversed[value]
	return key, ok
}

// Delete removes the key-value pair associated with the given key.
func (b *bimap[K, V]) Delete(key K) {
	if value, ok := b.forward[key]; ok {
		delete(b.forward, key)
		delete(b.reversed, value)
	}
}

// Len returns the number of key-value pairs in the BiMap.
func (b *bimap[K, V]) Len() int {
	return len(b.forward)
}

// Keys returns a slice of all keys in the BiMap.
func (b *bimap[K, V]) Keys() []K {
	keys := make([]K, 0, len(b.forward))
	for key := range b.forward {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice of all values in the BiMap.
func (b *bimap[K, V]) Values() []V {
	values := make([]V, 0, len(b.reversed))
	for value := range b.reversed {
		values = append(values, value)
	}
	return values
}

// Clear removes all key-value pairs from the BiMap.
func (b *bimap[K, V]) Clear() {
	b.forward = make(map[K]V)
	b.reversed = make(map[V]K)
}
