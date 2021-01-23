package container

//TestContainer is the interface for implementing container testing utilities.
type TestContainer interface {
	// Start creates and starts a new testing container.
	Start() error
	// Stop stops the container and clean it up.
	Stop() error
}
