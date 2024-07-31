package bimap

import (
	"slices"
	"testing"
)

func TestNewBiMap(t *testing.T) {
	// Empty BiMap
	emptyBM := NewBiMap[any, any](nil)

	if emptyBM.Len() != 0 {
		t.Errorf("Expected empty BiMap, got length %d", emptyBM.Len())
	}

	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	// BiMap with initial values
	nonEmptyBM := NewBiMap(m)
	// Keys and Values
	keys := nonEmptyBM.Keys()
	values := nonEmptyBM.Values()

	// Check length
	if nonEmptyBM.Len() != len(m) {
		t.Errorf("Expected BiMap length %d, got %d", len(m), nonEmptyBM.Len())
	}

	// Check entries
	for k, v := range m {

		if slices.Contains(keys, k) != true {
			t.Errorf("Expected key %s in BiMap", k)
		}

		if slices.Contains(values, v) != true {
			t.Errorf("Expected value %v in BiMap", v)
		}

		if val, ok := nonEmptyBM.GetByKey(k); !ok || val != v {
			t.Errorf("Expected value %v for key %s, got %v", v, k, val)
		}

		if val, ok := nonEmptyBM.GetByValue(v); !ok || val != k {
			t.Errorf("Expected key %s for value %v, got %v", k, v, val)
		}
	}

	// Add new entry
	nonEmptyBM.Put("three", 3)

	// Check length
	if nonEmptyBM.Len() != len(m) {
		t.Errorf("Expected BiMap length %d, got %d", len(m)+1, nonEmptyBM.Len())
	}

	// Delete entry
	nonEmptyBM.Delete("three")

	// Check length
	if nonEmptyBM.Len() != len(m) {
		t.Errorf("Expected BiMap length %d, got %d", len(m), nonEmptyBM.Len())
	}

	// Clear BiMap
	nonEmptyBM.Clear()

	// Check length
	if nonEmptyBM.Len() != 0 {
		t.Errorf("Expected empty BiMap, got length %d", nonEmptyBM.Len())
	}
}

func TestNewReadOnlyBiMap(t *testing.T) {

	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	rBiMap := NewReadOnlyBiMap(m)

	// Keys and Values
	keys := rBiMap.Keys()
	values := rBiMap.Values()

	// Check length
	if rBiMap.Len() != len(m) {
		t.Errorf("Expected ReadOnlyBiMap length %d, got %d", len(m), rBiMap.Len())
	}

	// Check entries
	for k, v := range m {

		if slices.Contains(keys, k) != true {
			t.Errorf("Expected key %s in BiMap", k)
		}

		if slices.Contains(values, v) != true {
			t.Errorf("Expected value %v in BiMap", v)
		}

		if val, ok := rBiMap.GetByKey(k); !ok || val != v {
			t.Errorf("Expected value %v for key %s, got %v", v, k, val)
		}

		if val, ok := rBiMap.GetByValue(v); !ok || val != k {
			t.Errorf("Expected key %s for value %v, got %v", k, v, val)
		}
	}
}
