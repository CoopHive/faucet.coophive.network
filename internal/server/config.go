package server

type Config struct {
	Network         string
	Symbol          string
	HttpPort        int
	Interval        int
	Payout          int
	TokenPayout     int
	ProxyCount      int
	HcaptchaSiteKey string
	HcaptchaSecret  string
}

func NewConfig(network, symbol string, httpPort, interval, payout, tokenPayout, proxyCount int, hcaptchaSiteKey, hcaptchaSecret string) *Config {
	return &Config{
		Network:         network,
		Symbol:          symbol,
		HttpPort:        httpPort,
		Interval:        interval,
		Payout:          payout,
		TokenPayout:     tokenPayout,
		ProxyCount:      proxyCount,
		HcaptchaSiteKey: hcaptchaSiteKey,
		HcaptchaSecret:  hcaptchaSecret,
	}
}
