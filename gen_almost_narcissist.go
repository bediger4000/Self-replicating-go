package main

import (
	"fmt"
	"os"
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
	b := make([]byte, 1)
	src := fmt.Sprintf(h, h)
	buffer := make([]byte, 0)
	crc := crc32.ChecksumIEEE([]byte(src))
	var xcrc uint32
	r := 0

	for i := 0; true; i++ {
		_, e := os.Stdin.Read(b)
		if e != nil {
			if e != io.EOF {
				r = 0
			} else {
				xcrc = crc32.ChecksumIEEE(buffer)
				if xcrc == crc {
					r = 1
				} else {
					r = 0
				}
			}
			break
		} else {
			buffer = append(buffer, b[0])
		}
	}

	fmt.Printf("%%d\n", r)
}
`
	src := fmt.Sprintf(h, h)
	fmt.Printf("%s", src)
}
