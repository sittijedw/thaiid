package main

import (
	"fmt"

	"github.com/sittijedw/thaiid/thaiid"
)

func main() {
	id := "1234567890121"

	fmt.Printf("Thaiid: %v, IsValid: %v", id, thaiid.IsValid(id))
}
