package filter

type OptionFunc func(*option)

type option struct {
	instreams  []string
	content    string
	outstreams []string
}

func InStream(s ...string) OptionFunc {
	return func(o *option) {
		o.instreams = append(o.instreams, s...)
	}
}

func Content(c string) OptionFunc {
	return func(o *option) {
		o.content = c
	}
}

func OutStream(s ...string) OptionFunc {
	return func(o *option) {
		o.outstreams = append(o.outstreams, s...)
	}
}
