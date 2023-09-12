package structures

import "testing"

func BenchmarkLookupTable_GetAPI(b *testing.B) {
	l := NewLookupTable()
	l.Append(Method{
		Identifier: "test",
	})
	for i := 0; i < b.N; i++ {
		l.GetAPI("test")
	}
}

func BenchmarkLookupTable_GetAPI_100(b *testing.B) {
	l := NewLookupTable()
	for i := 0; i < 100; i++ {
		l.Append(Method{
			Identifier: "test",
		})
	}
	for i := 0; i < b.N; i++ {
		l.GetAPI("test")
	}
}

func BenchmarkLookupTable_GetAPI_1000(b *testing.B) {
	l := NewLookupTable()
	for i := 0; i < 1000; i++ {
		l.Append(Method{
			Identifier: "test",
		})
	}
	for i := 0; i < b.N; i++ {
		l.GetAPI("test")
	}
}

func BenchmarkLookupTable_GetAPI_BAppend(b *testing.B) {
	l := NewLookupTable()
	for i := 0; i < b.N; i++ {
		l.Append(Method{
			Identifier: "test",
		})
	}
	for i := 0; i < b.N; i++ {
		l.GetAPI("test")
	}
}
