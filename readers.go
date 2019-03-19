package main

// Implement a Reader type that emits an infinit stream of the ASCII character 'A'.

// MyReader will need to satisfy the Reader interface, it firsts needs to implement a Read() method
type MyReader struct{}

// Add a Read method to MyReader -> Read([]byte) (int, err)
func (m MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}
