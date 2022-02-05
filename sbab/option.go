package xq

type options struct {
	source string
	mask   string
}

type OptFunc func(opts *options)

func SetSource(value string) OptFunc {
	return func(opts *options) {
		opts.source = value
	}
}

func SetMask(value string) OptFunc {
	return func(opts *options) {
		opts.mask = value
	}
}
