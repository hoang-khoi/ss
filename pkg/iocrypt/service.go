package iocrypt

import "io"

// Service takes io.Reader, does something and dump the result to io.Writer, returns the number of bytes read.
type Service interface {
	Process(dest io.Writer, src io.Reader, c chan<- int) error
}
