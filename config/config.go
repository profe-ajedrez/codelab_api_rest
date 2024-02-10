package config

import (
	"flag"
)

var cfg = newDefault()

type Config struct {
	Env   string
	Debug bool
	Mock  bool
	Port  int
}

func newDefault() *Config {
	return &Config{
		Env:   "development",
		Debug: false,
		Mock:  false,
		Port:  3000,
	}
}

func LoadCliConfig() {
	env := flag.String("env", "development", "environment")
	debug := flag.Bool("debug", false, "debug mode")
	mock := flag.Bool("mock", false, "mock database connection")
	port := flag.Int("port", -1, "port where to listen")

	flag.Parse()

	if env != nil && (*env == "production" || *env == "development") {
		cfg.Env = *env
	}

	cfg.Debug = *debug
	cfg.Mock = *mock

	if *port > -1 {
		cfg.Port = *port
	}
}

func Environment() string {
	return cfg.Env
}

func Debug() bool {
	return cfg.Debug
}

func Port() int {
	return cfg.Port
}

func Mock() bool {
	return cfg.Mock
}

func SetMockingOn() {
	cfg.Mock = true
}

func SetMockingOff() {
	cfg.Mock = false
}
