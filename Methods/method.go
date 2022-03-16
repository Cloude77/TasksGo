package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o") // arguments
	fixed := replacer.Replace(broken)         // replacement
	fmt.Println(fixed)
}
