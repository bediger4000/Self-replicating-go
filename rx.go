package main

import "fmt"

func main() {
	h := `package main

import "fmt"

func main() {
	h := %q

	fmt.Printf(h, h)
}
`
	fmt.Printf(h, h)
}
