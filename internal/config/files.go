package config

import (
	"os"
	"path/filepath"
)

const configDir = ".slog"

var (
	CAFile         = configFile("ca.pem")
	ServerCertFile = configFile("server.pem")
	ServerKeyFile  = configFile("server-key.pem")
	ClientCertFile = configFile("client.pem")
	ClientKeyfile  = configFile("client-key.pem")
)

func configFile(filename string) string {
	if dir := os.Getenv("SLOG_CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, configDir, filename)
}
