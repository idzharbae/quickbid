package config

import "fmt"

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

func (db Database) ToConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", db.Username, db.Password, db.Address, db.Port, db.DBName)
}
