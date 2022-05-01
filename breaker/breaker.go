package breaker

type (
	// is check the error can be acceptable
	Acceptable func(err error) bool

	// defind breaker interface
	Breaker interface {
		// return the breaker name
		Name() string

		// Allow() (Promise, error)
	}

	Promise interface {
		// tells breaker successful.
		Accept()
		// Reject tells the Breaker that the call is failed.
		Reject(reason string)
	}
)
