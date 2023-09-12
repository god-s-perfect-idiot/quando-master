package structures

import (
	"testing"
)

func BenchmarkMemoryInsert(b *testing.B) {
	db := NewMemory()
	for i := 0; i < b.N; i++ {
		db.Set("test", "test")
	}
}

func BenchmarkMemoryGet(b *testing.B) {
	db := NewMemory()
	db.Set("test", "test")
	for i := 0; i < b.N; i++ {
		db.Get("test")
	}
}

func BenchmarkMemoryDelete(b *testing.B) {
	db := NewMemory()
	db.Set("test", "test")
	for i := 0; i < b.N; i++ {
		db.Delete("test")
	}
}

func BenchmarkMemoryCheck(b *testing.B) {
	db := NewMemory()
	db.Set("test", "test")
	for i := 0; i < b.N; i++ {
		db.Check("test")
	}
}

func BenchmarkMemoryClear(b *testing.B) {
	db := NewMemory()
	db.Set("test", "test")
	for i := 0; i < b.N; i++ {
		db.Clear()
	}
}
