package main

import (
	"fmt"
)

func main() {
	h := `package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	h := %q
	var b [1024]byte
	var buffer []byte
	src := []byte(fmt.Sprintf(h, h))
	r := 0

	for {
		if n, e := os.Stdin.Read(b[0:]); n == 0 || e != nil {
			if bytes.Equal(src, buffer) {
				r = 1
			}
			break
		} else {
			buffer = append(buffer, b[0:n]...)
		}
		if len(buffer) > len(src) {
			break
		}
	}

	fmt.Printf("%%d\n", r)
}
`
	fmt.Printf(h, h)
}
