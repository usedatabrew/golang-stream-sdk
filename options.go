package streamdk

type Options struct {
	ApiKey     string
	StreamHost string
	ApiHost    string
}

func NewOptions(options ...Option) *Options {
	opts := &Options{
		StreamHost: "prod.gateway.databrew.tech:9000",
		ApiHost:    "https://api.databrew.tech",
	}

	for _, o := range options {
		o(opts)
	}
	return opts
}

type Option func(*Options)

func WithApiKey(apiKey string) Option {
	return func(o *Options) {
		o.ApiKey = apiKey
	}
}

func WithStreamHost(streamHost string) Option {
	return func(o *Options) {
		o.StreamHost = streamHost
	}
}

func WithApiHost(apiHost string) Option {
	return func(o *Options) {
		o.ApiHost = apiHost
	}
}
