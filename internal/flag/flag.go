package flag

import (
	"flag"
	"fmt"
)

var nFlag = flag.Int("n", 1234, "help message for flag n")

func flagParse() {
    fmt.Println("Flag: %d", nFlag)
}
