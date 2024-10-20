package config

type Config struct {
	CallbackPort       int
	CallbackPath       string
	ProxmoxURL         string
	TimeoutSeconds     int
	VerboseLog         bool
	Realm              string
	OutputFormat       string
	OpenDefaultBrowser bool
}
