package tournament

import (
	"bytes"
	"io"
)

func Tally(r io.Reader, w io.Writer) error {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(r)
	s := buf.String()
}
