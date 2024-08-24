package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

// App config struct
type Config struct {
	Server        ServerConfig
	Redis         RedisConfig
	RedisSentinel RedisSentinelConfig
	Logger        Logger
}

// Server config struct
type ServerConfig struct {
	AppVersion           string
	Port                 string
	DevMode              bool
	SSL                  bool
	CtxDefaultTimeout    time.Duration
	CSRF                 bool
	Debug                bool
	MaxConnectionIdle    time.Duration
	Timeout              time.Duration
	MaxConnectionAge     time.Duration
	ServiceName          string
	TimeConvert          string
	WaitShotDownDuration int
}

// Logger config
type Logger struct {
	DevMode  bool
	Encoder  string
	Encoding string
	LogLevel string
}

// Redis config
type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultdb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}
type RedisSentinelConfig struct {
	Addr             string
	RouteByLatency   bool // Allows routing read-only commands to the closest master or slave node.
	RouteRandomly    bool // Allows routing read-only commands to the random master or slave node.
	Username         string
	Password         string
	SentinelUsername string
	SentinelPassword string
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	PoolFIFO         bool // PoolFIFO uses FIFO mode for each node connection pool GET/PUT (default LIFO).
	PoolSize         int
	MinIdleConns     int
	MaxRetries       int
	MinRetryBackoff  time.Duration
	MaxRetryBackoff  time.Duration
	DialTimeout      time.Duration
	PoolTimeout      time.Duration
	MasterName       string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

// Get config
func InitConfig(env string) (*Config, error) {
	var configPath string
	switch env {
	case "qc":
		configPath = "./config/qc"
	case "staging":
		configPath = "./config/staging"
	case "prod":
		configPath = "./config/prod"
	default:
		configPath = "./config/local"
	}

	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
