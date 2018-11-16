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
)

func main() {
	h := %q
	b := make([]byte, 1)
	src := fmt.Sprintf(h, h)
	r := 0

	for i := 0; true; i++ {
		_, e := os.Stdin.Read(b)
		if e != nil {
			if e != io.EOF {
				r = 0
			} else {
				if i == len(src) {
					r = 1
				}
			}
			break
		}
		if i >= len(src) {
			r = 0
			break
		}
		if b[0] != src[i] {
			r = 0
			break
		}
	}

	fmt.Printf("%%d\n", r)
}
`
	src := fmt.Sprintf(h, h)
	fmt.Printf("%s", src)
}
