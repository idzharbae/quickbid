package config

type Config struct {
	DB               Database `json:"database"`
	ConnectionString string   `json:"connection_string"`
}

type Database struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}
