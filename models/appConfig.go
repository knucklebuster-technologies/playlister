package models

type serverConfig struct {
	Address string `json:"address"`
}

type dataConfig struct {
	URI    string `json:"uri"`
	DBName string `json:"dbname"`
}

// AppConfig structure to contain configuration values
type AppConfig struct {
	Version string
	Server  serverConfig `json:"server"`
	Data    dataConfig   `json:"data"`
}
