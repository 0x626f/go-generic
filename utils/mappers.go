package utils

// MapToKeySlice extracts all keys from a map and returns them as a slice.
// This is useful when you need to iterate over map keys in a specific order
// or perform operations on just the keys.
//
// Type parameters:
//   - K: The key type of the map must be comparable
//   - V: The value type of the map
//
// Parameters:
//   - m: The map from which to extract keys
//
// Returns a slice containing all the keys from the input map.
// Note: The order of keys in the returned slice is not deterministic.
func MapToKeySlice[K comparable, V any](m map[K]V) []K {
	slice := make([]K, len(m))

	for key, _ := range m {
		slice = append(slice, key)
	}
	return slice
}

// MapToValueSlice extracts all values from a map and returns them as a slice.
// This is useful when you need to iterate over map values in a specific order
// or perform operations on just the values.
//
// Type parameters:
//   - K: The key type of the map must be comparable
//   - V: The value type of the map
//
// Parameters:
//   - m: The map from which to extract values
//
// Returns a slice containing all the values from the input map.
// Note: The order of values in the returned slice is not deterministic
// and corresponds to the order of keys in the map's internal representation.
func MapToValueSlice[K comparable, V any](m map[K]V) []V {
	slice := make([]V, len(m))

	for _, value := range m {
		slice = append(slice, value)
	}
	return slice
}
