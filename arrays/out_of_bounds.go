package arrays

// This code won't compile because the compiler automatically detects
// invalid array indexes due to the very strong type system.
func out_of_bounds(array [5]string) {
	array[10] = ""
}
