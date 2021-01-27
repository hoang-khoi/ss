package strategy

import "io"

// Strategy defines the logic to be used for processing a block of data getting from io.Reader.
type Strategy interface {
	Process(dest io.Writer, src io.Reader) (int, error)
}
