package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	b64writer := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	h := `package main

import (
	"encoding/base64"
	"fmt"
	"os"
)


func main() {
	b64writer := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	h := %q

	b64writer.Write([]byte(fmt.Sprintf(h, h)))
}
`
	b64writer.Write([]byte(fmt.Sprintf(h, h)))
}
