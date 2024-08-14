package petal

type Client struct {
	AppKey      string
	AppSecret   string
	CallbackURL string
}

func NewClient(config *Config) (*Client, error) {
	client := &Client{
		AppKey:      config.AppKey,
		AppSecret:   config.AppSecret,
		CallbackURL: config.CallbackURL,
	}

	return client, nil
}

type Config struct {
	AppKey      string
	AppSecret   string
	CallbackURL string
}

func NewConfig(appKey, appSecret, callbackURl string) *Config {
	return &Config{AppKey: appKey, AppSecret: appSecret, CallbackURL: callbackURl}
}
