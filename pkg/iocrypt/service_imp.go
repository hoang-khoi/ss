package iocrypt

import (
	"io"
	"ss/pkg/iocrypt/strategy"
)

// ServiceImp implements Service with a simple strategy pattern.
type ServiceImp struct {
	Strategy strategy.Strategy
}

// Process uses Strategy to process io.Writer and dumps the result to io.Reader, it is done by block-by-block basis,
// for each loop, a number of processed bytes is passed to channel c.
func (s *ServiceImp) Process(dest io.Writer, src io.Reader, c chan<- int) error {
	for {
		n, err := s.Strategy.Process(dest, src)
		if err != nil {
			closeNullableChannel(c)
			if err != io.EOF {
				return err
			}
			break
		}

		writeNullableChannel(n, c)
	}

	return nil
}

func writeNullableChannel(n int, c chan<- int) {
	if c != nil {
		c <- n
	}
}

func closeNullableChannel(c chan<- int) {
	if c != nil {
		close(c)
	}
}
