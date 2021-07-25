package config

type AWS struct {
	Region string `toml:"region"`
	SQS    SQS
}

type SQS struct {
	BaseURL string `toml:"base_url"`
}
