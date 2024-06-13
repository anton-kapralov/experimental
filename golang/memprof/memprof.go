package main

func main() {
	createSlice()
	createLargeSlice()
	createSliceWithSize(1024)
}

//go:noinline
func createSlice() {
	_ = make([]int, 1024)
}

//go:noinline
func createLargeSlice() {
	_ = make([]int, 10_024)
}

//go:noinline
func createSliceWithSize(n int) {
	_ = make([]int, n)
}
