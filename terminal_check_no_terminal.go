// +build js nacl plan9 wasi

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
