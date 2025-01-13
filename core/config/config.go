package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	instance *Config
	once     sync.Once
	mu       sync.Mutex
)

func GetConfig(config_path string) *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig(config_path, instance); err != nil {
			panic(err)
		}
	})
	return instance
}

// inject custom config for testing
func SetConfig(cfg *Config) {
	mu.Lock()
	defer mu.Unlock()
	instance = cfg
}

// clear the injected config for testing
func ResetConfig() {
	mu.Lock()
	defer mu.Unlock()
	instance = nil
	once = sync.Once{} // reload once to allow loading config again
}
