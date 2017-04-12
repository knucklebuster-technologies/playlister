package models

// ServerConfig values for http server
type ServerConfig struct {
	Address string `json:"address"`
}

// DataConfig values for database server
type DataConfig struct {
	URI    string `json:"uri"`
	DBName string `json:"dbname"`
}

// DiscogsConfig values for the discogs web service
type DiscogsConfig struct {
	URI    string `json:"uri"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// AppConfig structure to contain configuration values
type AppConfig struct {
	Version string
	Server  ServerConfig  `json:"server"`
	Data    DataConfig    `json:"data"`
	Discogs DiscogsConfig `json:"discogs"`
}
