package main

import "testing"

func BenchmarkCreateSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createSlice()
	}
}

func BenchmarkCreateLargeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createLargeSlice()
	}
}

func BenchmarkCreateSliceWithSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createSliceWithSize(1024)
	}
}
