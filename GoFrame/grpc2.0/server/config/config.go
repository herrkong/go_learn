package config

import (
	
)

type ServerInfo struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	// CertFile string `yaml:"cert_file"`
	// KeyFile  string `yaml:"key_file"`
}

func GetServerConfig() *ServerInfo{
	return &ServerInfo{
		Host: "127.0.0.1",
		Port: 9601,
	}
}