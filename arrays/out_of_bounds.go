package arrays

// This code won't compile because the compiler automatically detects
// invalid array indexes because of the strong type system.
func out_of_bounds(array [5]string) {
	array[10] = ""
}
