// A wrapper of io.Writer returns a new Writer wraps the original
// and a pointer to an int64 variable
package main

// type wc struct {
// 	w       io.Writer
// 	written int64
// }
//
// func (c *wc) Write(p []byte) (int, error) {
// 	c.written += int64(len(p))
// 	return len(p), nil
// }
//
// // returns a new Writer wraps the original and a pointer to an int64
// func CountingWriter(w io.Writer) (io.Writer, *int64) {
// 	var wout wc
//
// 	return w, &c
// }

func main() {

}
