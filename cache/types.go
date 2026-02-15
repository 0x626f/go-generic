package cache

type record[K comparable, V any] struct {
	key   K
	value V
}
