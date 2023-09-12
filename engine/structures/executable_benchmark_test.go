package structures

import "testing"

func BenchmarkExecutable_GetData(b *testing.B) {
	e := Executable{
		Data: make(map[string]interface{}),
	}
	e.SetData("test", "test")
	for i := 0; i < b.N; i++ {
		e.GetData("test")
	}
}

func BenchmarkExecutable_SetData(b *testing.B) {
	e := Executable{
		Data: make(map[string]interface{}),
	}
	for i := 0; i < b.N; i++ {
		e.SetData("test", "test")
	}
}

func BenchmarkExecutable_ValidateData(b *testing.B) {
	e := Executable{
		Data: make(map[string]interface{}),
	}
	e.SetData("test", "test")
	for i := 0; i < b.N; i++ {
		e.ValidateData("test")
	}
}

func BenchmarkExecutable_CheckVal(b *testing.B) {
	e := Executable{
		Val: 1.0,
	}
	for i := 0; i < b.N; i++ {
		e.CheckVal()
	}
}

func BenchmarkExecutable_GetData_100(b *testing.B) {
	e := Executable{
		Data: make(map[string]interface{}),
	}
	for i := 0; i < 100; i++ {
		e.SetData("test", "test")
	}
	for i := 0; i < b.N; i++ {
		e.GetData("test")
	}
}

func BenchmarkExecutable_SetData_100(b *testing.B) {
	e := Executable{
		Data: make(map[string]interface{}),
	}
	for i := 0; i < b.N; i++ {
		e.SetData("test", "test")
	}
}

func BenchmarkExecutable_ValidateData_100(b *testing.B) {
	e := Executable{
		Data: make(map[string]interface{}),
	}
	for i := 0; i < 100; i++ {
		e.SetData("test", "test")
	}
	for i := 0; i < b.N; i++ {
		e.ValidateData("test")
	}
}

func BenchmarkExecutable_CheckVal_100(b *testing.B) {
	e := Executable{
		Val: 1.0,
	}
	for i := 0; i < b.N; i++ {
		e.CheckVal()
	}
}

func BenchmarkExecutable_GetData_B(b *testing.B) {
	e := Executable{
		Data: make(map[string]interface{}),
	}
	for i := 0; i < b.N; i++ {
		e.SetData("test", "test")
	}
	for i := 0; i < b.N; i++ {
		e.GetData("test")
	}
}
