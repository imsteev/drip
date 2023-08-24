package utils

import (
	"fmt"
	"io"
)

func WriteStrf(w io.Writer, s string, args ...any) {
	w.Write([]byte(fmt.Sprintf(s, args...)))
}
