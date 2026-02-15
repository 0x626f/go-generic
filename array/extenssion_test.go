package array

import (
	"math/rand"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	size := 1_000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := New[int]()

	for range size {
		rnd := r.Intn(10000)
		array.Push(rnd)
	}

	array.InsertionSort(func(arg0, arg1 int) int {
		if arg0 > arg1 {
			return 1
		} else if arg0 < arg1 {
			return -1
		}
		return 0
	})

	for index := 1; index < size; index++ {
		if array.At(index-1) > array.At(index) {
			t.Fatal("wrong insertion sort")
		}
	}
}

func TestHeapSort(t *testing.T) {
	size := 1_000_000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := New[int]()

	for range size {
		rnd := r.Intn(10000)
		array.Push(rnd)
	}

	array.HeapSort(func(arg0, arg1 int) int {
		if arg0 > arg1 {
			return 1
		} else if arg0 < arg1 {
			return -1
		}
		return 0
	})

	for index := 1; index < size; index++ {
		if array.At(index-1) > array.At(index) {
			t.Fatal("wrong insertion sort")
		}
	}
}

func TestBinarySearch(t *testing.T) {
	size := 1_000_000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := New[int]()

	for range size {
		rnd := r.Intn(10000)
		array.Push(rnd)
	}

	array.HeapSort(func(arg0, arg1 int) int {
		if arg0 > arg1 {
			return 1
		} else if arg0 < arg1 {
			return -1
		}
		return 0
	})

	target := array.At(57)

	result, found := array.BinarySearch(target, func(arg0, arg1 int) int {
		if arg0 > arg1 {
			return 1
		} else if arg0 < arg1 {
			return -1
		}
		return 0
	})

	if !found {
		t.Fatal("not found")
	}

	if result != target {
		t.Fatal("wrong search")
	}
}
func BenchmarkInsertionSort(b *testing.B) {
	size := 1_000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := New[int]()

	for range size {
		rnd := r.Intn(10000)
		array.Push(rnd)
	}

	b.StartTimer()
	array.InsertionSort(func(arg0, arg1 int) int {
		if arg0 > arg1 {
			return 1
		} else if arg0 < arg1 {
			return -1
		}
		return 0
	})
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkHeapSort(b *testing.B) {
	size := 1_000_000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := New[int]()

	for range size {
		rnd := r.Intn(10000)
		array.Push(rnd)
	}

	b.StartTimer()
	array.HeapSort(func(arg0, arg1 int) int {
		if arg0 > arg1 {
			return 1
		} else if arg0 < arg1 {
			return -1
		}
		return 0
	})
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkBinarySearch(b *testing.B) {
	size := 1_000_000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := New[int]()

	for range size {
		rnd := r.Intn(10000)
		array.Push(rnd)
	}

	array.HeapSort(func(arg0, arg1 int) int {
		if arg0 > arg1 {
			return 1
		} else if arg0 < arg1 {
			return -1
		}
		return 0
	})

	target := array.At(57)

	b.StartTimer()
	result, found := array.BinarySearch(target, func(arg0, arg1 int) int {
		if arg0 > arg1 {
			return 1
		} else if arg0 < arg1 {
			return -1
		}
		return 0
	})
	b.StopTimer()
	b.ReportAllocs()

	if !found {
		b.Fatal("not found")
	}

	if result != target {
		b.Fatal("wrong search")
	}
}
