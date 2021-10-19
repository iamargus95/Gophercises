package camelcase

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	result := camelCase(flag.Arg(0))
	fmt.Printf("The string %s has %d words.\n", flag.Arg(0), result)
}

func camelCase(input string) int {

}
