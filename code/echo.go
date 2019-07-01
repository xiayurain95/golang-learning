package code

import (
	"fmt"
	"os"
)
// Echo os.Args
func Echo() {
	for i := range os.Args {
		fmt.Printf(string(i))
	}
}
