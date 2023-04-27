package featurelab

import "time"

type Config struct {
	cacheTTL time.Duration
}

var DefaultConfig = Config{
	cacheTTL: 60 * time.Minute,
}
