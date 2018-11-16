package main

import (
	"fmt"
)

func main() {
	h := `package main

import (
	"fmt"
	"io"
	"os"
	"hash/crc32"
)

func main() {
	h := %q
	var b [1024]byte
	var buffer []byte
	r := 0
	for {
		if n, e := os.Stdin.Read(b[:]); n == 0 || e != nil {
			if n == 0 || e == io.EOF {
				if crc32.ChecksumIEEE([]byte(fmt.Sprintf(h, h))) == crc32.ChecksumIEEE(buffer) {
					r = 1
				}
			}
			break
		} else {
			buffer = append(buffer, b[0:n]...)
		}
	}
	fmt.Printf("%%d\n", r)
}
`
	fmt.Printf(h, h)
}
