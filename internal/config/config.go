package config

import (
	"fmt"
	"net/url"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"github.com/vadimklimov/cpi-mcp-server/internal/util/logger"
)

type Config struct {
	BaseURL        *url.URL        `mapstructure:"base_url"`
	TokenURL       *url.URL        `mapstructure:"token_url"`
	ClientID       string          `mapstructure:"client_id"`
	ClientSecret   string          `mapstructure:"client_secret"`
	MaxConcurrency int             `mapstructure:"max_concurrency"`
	Timeout        int             `mapstructure:"timeout"`
	Transport      Transport       `mapstructure:"transport"`
	Port           string          `mapstructure:"port"`
	LogLevel       logger.LogLevel `mapstructure:"log_level"`
	LogFile        string          `mapstructure:"log_file"`
}

var (
	instance *Config
	once     sync.Once
)

type Transport string

const (
	TransportStdio Transport = "stdio"
	TransportHTTP  Transport = "http"
)

func Init() error {
	var initError error

	once.Do(func() {
		supportedConfigParams := []string{
			"base_url",
			"token_url",
			"client_id",
			"client_secret",
			"max_concurrency",
			"timeout",
			"transport",
			"port",
			"log_level",
			"log_file",
		}

		viper.SetEnvPrefix("mcp_cpi")
		viper.AutomaticEnv()

		for _, configParam := range supportedConfigParams {
			if err := viper.BindEnv(configParam); err != nil {
				initError = fmt.Errorf("failed to bind environment variable '%s': %v", configParam, err)

				return
			}
		}

		if err := viper.Unmarshal(&instance, viper.DecodeHook(composeDecodeHook())); err != nil {
			initError = fmt.Errorf("failed to unmarshal configuration: %v", err)

			return
		}

		if err := instance.checkRequired(); err != nil {
			initError = fmt.Errorf("required configuration parameters were not provided: %v", err)

			return
		}

		instance.setDefaults()
	})

	return initError
}

func BaseURL() *url.URL {
	return instance.BaseURL
}

func TokenURL() *url.URL {
	return instance.TokenURL
}

func ClientID() string {
	return instance.ClientID
}

func ClientSecret() string {
	return instance.ClientSecret
}

func MaxConcurrency() int {
	return instance.MaxConcurrency
}

func Timeout() int {
	return instance.Timeout
}

func ServerTransport() Transport {
	return instance.Transport
}

func ServerPort() string {
	return instance.Port
}

func LogFile() string {
	return instance.LogFile
}

func LogLevel() logger.LogLevel {
	return instance.LogLevel
}

func (c *Config) checkRequired() error {
	missingConfigParams := make([]string, 0)

	if c.BaseURL == nil {
		missingConfigParams = append(missingConfigParams, "base_url")
	}

	if c.TokenURL == nil {
		missingConfigParams = append(missingConfigParams, "token_url")
	}

	if c.ClientID == "" {
		missingConfigParams = append(missingConfigParams, "client_id")
	}

	if c.ClientSecret == "" {
		missingConfigParams = append(missingConfigParams, "client_secret")
	}

	if len(missingConfigParams) > 0 {
		return fmt.Errorf("missing parameters: %s", strings.Join(missingConfigParams, ", "))
	}

	return nil
}

func (c *Config) setDefaults() {
	supportedTransports := []Transport{
		TransportStdio,
		TransportHTTP,
	}

	supportedLogLevels := []logger.LogLevel{
		logger.LogLevelNone,
		logger.LogLevelError,
		logger.LogLevelWarn,
		logger.LogLevelInfo,
		logger.LogLevelDebug,
	}

	const (
		defaultTimeout   = 60
		defaultTransport = TransportStdio
		defaultPort      = "8080"
		defaultLogLevel  = logger.LogLevelNone
		defaultLogFile   = "/var/log/cpi-mcp-server.log"
	)

	DefaultMaxConcurrency := runtime.NumCPU()

	if c.MaxConcurrency <= 0 {
		c.MaxConcurrency = DefaultMaxConcurrency
	}

	if c.Timeout <= 0 {
		c.Timeout = defaultTimeout
	}

	if !slices.Contains(supportedTransports, c.Transport) {
		c.Transport = defaultTransport
	}

	if c.Port == "" {
		c.Port = defaultPort
	} else {
		if port, err := strconv.Atoi(c.Port); err != nil || port < 1024 || port > 65535 {
			c.Port = defaultPort
		}
	}

	if !slices.Contains(supportedLogLevels, c.LogLevel) {
		c.LogLevel = defaultLogLevel
	}

	if c.LogFile == "" {
		c.LogFile = defaultLogFile
	}
}
