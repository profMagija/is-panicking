package ispanicking

import "C"

//go:linkname IsPanicking is_panicking
func IsPanicking() bool
